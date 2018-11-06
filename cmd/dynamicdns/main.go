package main

import (
	"log"
	"time"

	"github.com/t0k4rt/dynamic-dns/internal/dnsupdater"
	"github.com/t0k4rt/dynamic-dns/internal/dnsupdater/gandi"
	"github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider"
	"github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider/livebox"
)

func main() {

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
