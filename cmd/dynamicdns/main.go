package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/t0k4rt/dynamic-dns/models/config"
)

func main() {
	cobra.OnInitialize()

	var config = config.NewTomlConfig()

	if _, err := toml.DecodeFile("config/example.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	log.Println(config.General)
	log.Println(config.Domain)
	forever := make(chan bool)

	for _, dom := range config.Domain {
		updater(dom, config.General)
	}
	<-forever
}

func updater(domain config.Domain, general config.General) {

	go func() {

		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {

			currentIP, err := domain.IPProvider.GetIP()
			if err != nil {
				log.Fatalln(err)
			}

			err = domain.DNSProvider.UpdateDNS(domain.Name.URL, currentIP)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("DNS updated !")
		}
	}()
}
