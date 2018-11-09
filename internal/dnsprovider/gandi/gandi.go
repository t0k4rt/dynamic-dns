package gandi

import (
	"errors"
	"net"
	"net/url"
	"os"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider"
	gandiApi "github.com/t0k4rt/gandi-livedns-go/client"
	"github.com/t0k4rt/gandi-livedns-go/client/domains"
	"github.com/t0k4rt/gandi-livedns-go/models"
)

type gandiDNSUpdater struct {
	gandiClient *gandiApi.GandiLiveDNS
	gandiAuth   runtime.ClientAuthInfoWriter
}

func NewDNSProvider() (*gandiDNSUpdater, error) {
	// check api key presence
	apiKey, ok := os.LookupEnv("GANDI_KEY")
	if !ok {
		return nil, errors.New("Gandi api key is missing, please set GANDI_KEY env var")
	}

	return &gandiDNSUpdater{
		gandiClient: gandiApi.Default,
		gandiAuth:   httptransport.APIKeyAuth("X-Api-Key", "header", apiKey),
	}, nil
}

func (l *gandiDNSUpdater) UpdateDNS(domain *url.URL, ip *ipprovider.ProvidedIP) error {
	err := l.update(domain, ip)
	if err != nil {
		return err
	}

	err = l.verifyIPV4(domain, ip.GetIPV4())
	if err != nil {
		return err
	}

	err = l.verifyIPV6(domain, ip.GetIPV6())
	if err != nil {
		return err
	}

	return nil
}

func (l *gandiDNSUpdater) update(domain *url.URL, ip *ipprovider.ProvidedIP) error {

	domainRecords := domains.NewPutDomainsDomainRecordsRecordNameParams()
	domainRecords.SetRecordName("@")
	domainRecords.SetDomain(domain.Hostname())
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
	}
	return nil
}

func (l *gandiDNSUpdater) verifyIPV4(domain *url.URL, ip net.IP) error {
	if ip == nil {
		return nil
	}

	p := domains.NewGetDomainsDomainRecordsRecordNameRecordTypeParams()
	p.SetDomain(domain.Hostname())
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

func (l *gandiDNSUpdater) verifyIPV6(domain *url.URL, ip net.IP) error {
	if ip == nil {
		return nil
	}

	p := domains.NewGetDomainsDomainRecordsRecordNameRecordTypeParams()
	p.SetDomain(domain.Hostname())
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