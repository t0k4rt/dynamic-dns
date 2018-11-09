package cmd

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/t0k4rt/dynamic-dns/models/config"
)

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
		log.Printf("Start updating dns with domain %s, delay %s, ttl %d", domain.Name.String(), refreshDelay, ttl)

		ticker := time.NewTicker(refreshDelay)
		for range ticker.C {
			currentIP, err := domain.IPProvider.GetIP()
			if err != nil {
				log.Fatalln(err)
			}

			err = domain.DNSProvider.UpdateDNS(domain.Name.URL, currentIP, ttl)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("DNS updated !")
		}
	}()
}
