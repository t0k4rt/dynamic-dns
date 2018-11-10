package cmd

import (
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	kitlog "github.com/go-kit/kit/log"
	kitlevel "github.com/go-kit/kit/log/level"
	"github.com/spf13/cobra"
	"github.com/t0k4rt/dynamic-dns/models/config"
)

var logger kitlog.Logger
var errLogger = log.New(os.Stderr, "", 0)

var cmdStart = &cobra.Command{
	Use:   "start",
	Short: "Start dns updater process",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startCmdRun(cfgFile)
	},
}

func init() {
	cmdStart.Flags().StringVar(&cfgFile, "config", "", "config file path")
	cmdStart.MarkFlagRequired("config")
}

func startCmdRun(cfgFile string) error {
	cfg, err := loadConfig(cfgFile)
	if err != nil {
		return err
	}
	initLogger(cfg.General.LogPath, cfg.General.LogLevel)
	kitlevel.Info(logger).Log("Event", "Dynamic dns started")
	for _, dom := range cfg.Domain {
		updater(dom, cfg.General)
	}

	forever := make(chan bool)
	<-forever
	return nil
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

func updater(domain config.Domain, general config.General) {

	go func() {
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
		kitlevel.Info(logger).Log("Event", "start updater", "Domain", domain.Name.String(), "refresh_delay", refreshDelay, "ttl", ttl)

		ticker := time.NewTicker(refreshDelay)
		for range ticker.C {
			currentIP, err := domain.IPProvider.GetIP()
			if err != nil {
				errLogger.Fatalln(err)
			}

			err = domain.DNSProvider.UpdateDNS(domain.Name.URL, currentIP, ttl)
			if err != nil {
				errLogger.Fatalln(err)
			}
			kitlevel.Info(logger).Log("Event", "domain updated", "domain", domain.Name.String())
		}
	}()
}
