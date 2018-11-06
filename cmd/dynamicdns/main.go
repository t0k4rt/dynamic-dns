package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/t0k4rt/dynamic-dns/internal/dnsupdater"
	"github.com/t0k4rt/dynamic-dns/internal/dnsupdater/gandi"
	"github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider"
	"github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider/livebox"
	"github.com/t0k4rt/dynamic-dns/models"
)

func main() {
	cobra.OnInitialize(initConfig)

	var config = models.NewTomlConfig()

	if _, err := toml.DecodeFile("config/example.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	log.Println(config.General)
	log.Println(config.Domain)

	forever := make(chan bool)

	liveboxProvider := livebox.NewProvider()

	gandiUpdater, err := gandi.NewUpdater("toktok.fr")
	if err != nil {
		log.Fatalln(err)
	}
	updater(liveboxProvider, gandiUpdater)
	<-forever
}

func updater(ipProvider ipaddressprovider.IPProvider, dnsProvider dnsupdater.DNSProvider) {

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {

			currentIP, err := ipProvider.GetIP()
			if err != nil {
				log.Fatalln(err)
			}

			err = dnsProvider.UpdateDNS(currentIP)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("DNS updated !")
		}
	}()
}
