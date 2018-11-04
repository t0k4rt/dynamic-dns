package gandi

import (
	"errors"
	"net"
	"net/url"
	"os"

	"github.com/go-openapi/runtime"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/t0k4rt/dynamic-dns/internal/ipaddressprovider"
	gandiApi "github.com/t0k4rt/gandi-livedns-go/client"
	"github.com/t0k4rt/gandi-livedns-go/client/domains"
	"github.com/t0k4rt/gandi-livedns-go/models"
)

type gandiDNSUpdater struct {
	domain      string
	gandiClient *gandiApi.GandiLiveDNS
	gandiAuth   runtime.ClientAuthInfoWriter
}

func NewUpdater(domain string) (*gandiDNSUpdater, error) {
	// validate url
	_, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	// check api key presence
	apiKey, ok := os.LookupEnv("GANDI_KEY")
	if !ok {
		return nil, errors.New("Gandi api key is missing, please set GANDI_KEY env var")
	}

	return &gandiDNSUpdater{
		domain:      domain,
		gandiClient: gandiApi.Default,
		gandiAuth:   httptransport.APIKeyAuth("X-Api-Key", "header", apiKey),
	}, nil
}

func (l *gandiDNSUpdater) UpdateDNS(ip *ipaddressprovider.ProvidedIP) error {
	err := l.update(ip)
	if err != nil {
		return err
	}

	err = l.verifyIPV4(ip.GetIPV4())
	if err != nil {
		return err
	}

	err = l.verifyIPV6(ip.GetIPV6())
	if err != nil {
		return err
	}

	return nil
}

func (l *gandiDNSUpdater) update(ip *ipaddressprovider.ProvidedIP) error {

	domainRecords := domains.NewPutDomainsDomainRecordsRecordNameParams()
	domainRecords.SetRecordName("@")
	domainRecords.SetDomain(l.domain)
	var records []*models.Record

	if ip.GetIPV4() != nil {
		records = append(records, &models.Record{
			RrsetName:   "@",
			RrsetTTL:    300,
			RrsetType:   "A",
			RrsetValues: []string{ip.GetIPV4String()},
		})
	}

	if ip.GetIPV6() != nil {
		records = append(records, &models.Record{
			RrsetName:   "@",
			RrsetTTL:    300,
			RrsetType:   "AAAA",
			RrsetValues: []string{ip.GetIPV6String()},
		})
	}

	domainRecords.SetRecord(domains.PutDomainsDomainRecordsRecordNameBody{
		Items: records,
	})

	_, err := l.gandiClient.Domains.PutDomainsDomainRecordsRecordName(domainRecords, l.gandiAuth)
	if err != nil {
		return err
		// switch e := err.(type) {
		// case *domains.PutDomainsDomainRecordsRecordNameBadRequest:
		// 	fmt.Printf("plop %v \n", e.Payload.Status)
		// case *domains.PutDomainsDomainRecordsRecordNameDefault:
		// 	fmt.Printf("plop %v \n", e.Payload.Message)
		// }
		// log.Fatal(err)
	}
	return nil
}

func (l *gandiDNSUpdater) verifyIPV4(ip net.IP) error {
	if ip == nil {
		return nil
	}

	p := domains.NewGetDomainsDomainRecordsRecordNameRecordTypeParams()
	p.SetDomain(l.domain)
	p.SetRecordName("@")
	p.SetRecordType("A")

	domainResp, err := l.gandiClient.Domains.GetDomainsDomainRecordsRecordNameRecordType(p, l.gandiAuth)

	if err != nil {
		return err
	}
	if domainResp.Payload.RrsetValues[0] == ip.String() {
		return nil
	}
	return nil
}

func (l *gandiDNSUpdater) verifyIPV6(ip net.IP) error {
	if ip == nil {
		return nil
	}

	p := domains.NewGetDomainsDomainRecordsRecordNameRecordTypeParams()
	p.SetDomain(l.domain)
	p.SetRecordName("@")
	p.SetRecordType("AAAA")

	domainResp, err := l.gandiClient.Domains.GetDomainsDomainRecordsRecordNameRecordType(p, l.gandiAuth)

	if err != nil {
		return err
	}
	if domainResp.Payload.RrsetValues[0] == ip.String() {
		return nil
	}
	return nil
}
