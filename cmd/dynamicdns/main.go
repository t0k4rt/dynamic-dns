package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	"github.com/spf13/cobra"
	"github.com/t0k4rt/dynamic-dns/models/config"
)

var cfgFile string
var logger kitlog.Logger
var errLogger = log.New(os.Stderr, "", 0)
var wg sync.WaitGroup

var rootCmd = &cobra.Command{
	Use:   "dynamicdns",
	Short: "Dynamic dns is a dyndns replacement",
}

var cmdStart = &cobra.Command{
	Use:   "start",
	Short: "Start dns updater process",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startCmdRun(cfgFile)
	},
}

func init() {
	rootCmd.AddCommand(cmdStart)
	cmdStart.Flags().StringVar(&cfgFile, "config", "", "config file path")
	cmdStart.MarkFlagRequired("config")
}

func monitorSignal(cancel context.CancelFunc) {
	// setup signal catching
	sigs := make(chan os.Signal, 1)
	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	go func() {
		for {
			select {
			case <-sigs:
				cancel()
				return
			}
		}
	}()
}

func startCmdRun(cfgFile string) error {
	cfg, err := loadConfig(cfgFile)
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	monitorSignal(cancel)
	initLogger(cfg.General.LogPath, cfg.General.LogLevel)
	kitlevel.Info(logger).Log("Event", "Dynamic dns started")

	for _, dom := range cfg.Domain {
		updater(ctx, dom, cfg.General)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}
}

func loadConfig(path string) (config.TomlConfig, error) {
	var config = config.NewTomlConfig()
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return config, err
	}
	return config, nil
}

func initLogger(path string, level string) {
	var f *os.File
	switch {
	case path == "stdout":
		f = os.Stdout
	default:
		var err error
		f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			errLogger.Fatalf("Could not write logs to %s", path)
		}
	}

	//defer f.Close()
	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(f))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	switch {
	case level == "debug":
		logger = kitlevel.NewFilter(logger, kitlevel.AllowDebug())
	case level == "info":
		logger = kitlevel.NewFilter(logger, kitlevel.AllowInfo())
	case level == "warn":
		logger = kitlevel.NewFilter(logger, kitlevel.AllowWarn())
	case level == "error":
		logger = kitlevel.NewFilter(logger, kitlevel.AllowError())
	case true:
		kitlevel.Warn(logger).Log("unknown log level %s, default to info", level)
	default:
		logger = kitlevel.NewFilter(logger, kitlevel.AllowInfo())
	}

}

func updater(ctx context.Context, domain config.Domain, general config.General) {
	wg.Add(1)
	go func(wg *sync.WaitGroup, ctx context.Context) {
		var refreshDelay time.Duration
		if domain.RefreshDelay.Seconds() == float64(0) {
			refreshDelay = general.DefaultRefreshDelay.Duration
		} else {
			refreshDelay = domain.RefreshDelay.Duration
		}
		var ttl int
		if domain.TTL != 0 {
			ttl = domain.TTL
		} else {
			ttl = general.DefaultTTL
		}

		var ipVersion int
		if domain.IPVersion == 0 {
			ipVersion = 4
		} else {
			ipVersion = domain.IPVersion
		}

		kitlevel.Info(logger).Log("Event", "start updater", "Domain", domain.Name, "refresh_delay", refreshDelay, "ttl", ttl, "ipversion", ipVersion)

		ticker := time.NewTicker(refreshDelay)
		for {
			select {
			case <-ticker.C:
				currentIP, err := domain.IPProvider.GetIP(ipVersion)
				if err != nil {
					kitlevel.Info(logger).Log("Error", "failed to get ip", "error", err.Error())
					errLogger.Fatalln(err)
				}
				kitlevel.Info(logger).Log("Event", "Ip retrieved", "ip", currentIP.String())

				err = domain.DNSProvider.UpdateDNS(domain.Name, currentIP, ttl, ipVersion)
				if err != nil {
					kitlevel.Info(logger).Log("Error", "failed to update dns", "error", err.Error())
					errLogger.Fatalln(err)
				}
				kitlevel.Info(logger).Log("Event", "domain updated", "domain", domain.Name)

			case <-ctx.Done():
				kitlevel.Info(logger).Log("Event", "Stopping updater", "Domain", domain.Name, "refresh_delay", refreshDelay, "ttl", ttl)
				wg.Done()
				return
			}
		}
	}(&wg, ctx)
}

func main() {
	done := make(chan bool, 1)
	go func() {
		if err := rootCmd.Execute(); err != nil {
			os.Exit(1)
		}
		wg.Wait()
		done <- true
	}()
	<-done
}
