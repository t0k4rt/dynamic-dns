package main

import (
	"log"
	"time"

	"github.com/t0k4rt/dynamic-dns/internal/dnsupdater/gandi"
	"github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider/livebox"
)

// 1xWDwzc9IZEm8ic01YKFZT7c

func main() {

	forever := make(chan bool)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {

			liveboxProvider := livebox.NewProvider()

			gandiUpdater, err := gandi.NewUpdater("toktok.fr")
			if err != nil {
				log.Fatalln(err)
			}

			currentIP, err := liveboxProvider.GetIP()
			if err != nil {
				log.Fatalln(err)
			}

			err = gandiUpdater.UpdateDNS(currentIP)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("DNS updated !")
		}
	}()
	<-forever
}

// func GetIP() *ip_address.PostSysbusNMCGetWANStatusOKBodyResultData {
// 	fmt.Println("retrieving IP")
// 	// create the API client, with the transport
// 	client := liveboxApi.Default

// 	// to override the host for the default client
// 	// apiclient.Default.SetTransport(transport)

// 	// make the request to get all items
// 	resp, err := client.IPAddress.PostSysbusNMCGetWANStatus(ip_address.NewPostSysbusNMCGetWANStatusParams())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("ip v4: %#v\n", resp.Payload.Result.Data.IPAddress)
// 	fmt.Printf("ip v6: %#v\n", resp.Payload.Result.Data.IPV6Address)
// 	return resp.Payload.Result.Data
// }

// func UpdateIP(newIP string) {
// 	fmt.Println("updating IP", newIP)
// 	// create the API client
// 	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Api-Key", "header", "1xWDwzc9IZEm8ic01YKFZT7c")

// 	// create transport and client
// 	gandiClient := gandiApi.Default

// 	domainRecords := domains.NewPutDomainsDomainRecordsRecordNameParams()
// 	domainRecords.SetRecordName("@")
// 	domainRecords.SetDomain("toktok.fr")

// 	record := models.Record{
// 		RrsetName:   "@",
// 		RrsetTTL:    300,
// 		RrsetType:   "A",
// 		RrsetValues: []string{newIP},
// 	}
// 	domainRecords.SetRecord(domains.PutDomainsDomainRecordsRecordNameBody{
// 		Items: []*models.Record{&record},
// 	})

// 	domainPutResp, err := gandiClient.Domains.PutDomainsDomainRecordsRecordName(domainRecords, apiKeyHeaderAuth)
// 	if err != nil {
// 		switch e := err.(type) {
// 		case *domains.PutDomainsDomainRecordsRecordNameBadRequest:
// 			fmt.Printf("plop %v \n", e.Payload.Status)
// 		case *domains.PutDomainsDomainRecordsRecordNameDefault:
// 			fmt.Printf("plop %v \n", e.Payload.Message)
// 		}
// 		log.Fatal(err)
// 	}
// 	fmt.Printf(">>> plop %v \n", domainPutResp.Payload.Message)

// 	p := domains.NewGetDomainsDomainRecordsRecordNameRecordTypeParams()
// 	p.SetDomain("toktok.fr")
// 	p.SetRecordName("@")
// 	p.SetRecordType("A")

// 	domainResp, err := gandiClient.Domains.GetDomainsDomainRecordsRecordNameRecordType(p, apiKeyHeaderAuth)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%#v\n", domainResp.Payload)
// 	if domainResp.Payload.RrsetValues[0] == newIP {
// 		fmt.Printf("GANDI updated")
// 	}
// }
