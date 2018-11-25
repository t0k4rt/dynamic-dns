package gandi

import (
	"errors"
	"net"
	"os"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/t0k4rt/dynamic-dns/pkg/domainparser"
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

func (l *gandiDNSUpdater) UpdateDNS(fullDomain string, ip net.IP, ttl int, version int) error {
	domain, recordName, err := domainparser.ParseDomain(fullDomain)
	if err != nil {
		return err
	}

	err = l.update(domain, recordName, ip, int32(ttl), version)
	if err != nil {
		return err
	}

	err = l.verifyIP(domain, recordName, ip, version)
	if err != nil {
		return err
	}

	return nil
}

func (l *gandiDNSUpdater) update(domain string, recordName string, ip net.IP, ttl int32, version int) error {

	domainRecords := domains.NewPutDomainsDomainRecordsRecordNameRecordTypeParams()
	domainRecords.SetRecordName(recordName)
	domainRecords.SetDomain(domain)
	var record *models.Record
	switch {
	case version == 4:
		domainRecords.SetRecordType("A")
		record = &models.Record{
			RrsetName:   recordName,
			RrsetTTL:    ttl,
			RrsetType:   "A",
			RrsetValues: []string{ip.String()},
		}
	case version == 6:
		domainRecords.SetRecordType("AAAA")
		record = &models.Record{
			RrsetName:   recordName,
			RrsetTTL:    ttl,
			RrsetType:   "AAAA",
			RrsetValues: []string{ip.String()},
		}
	}

	domainRecords.SetRecord(record)

	_, err := l.gandiClient.Domains.PutDomainsDomainRecordsRecordNameRecordType(domainRecords, l.gandiAuth)
	if err != nil {
		return err
	}
	return nil
}

func (l *gandiDNSUpdater) verifyIP(domain string, recordName string, ip net.IP, version int) error {
	if ip == nil {
		return nil
	}

	p := domains.NewGetDomainsDomainRecordsRecordNameRecordTypeParams()
	p.SetDomain(domain)
	p.SetRecordName(recordName)
	switch {
	case version == 4:
		p.SetRecordType("A")
	case version == 6:
		p.SetRecordType("AAAA")
	}

	domainResp, err := l.gandiClient.Domains.GetDomainsDomainRecordsRecordNameRecordType(p, l.gandiAuth)

	if err != nil {
		return err
	}
	if domainResp.Payload.RrsetValues[0] == ip.String() {
		return nil
	}
	return nil
}
