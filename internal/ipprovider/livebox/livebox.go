package livebox

import (
	"github.com/t0k4rt/dynamic-dns/internal/ipprovider"
	liveboxApi "github.com/t0k4rt/livebox-go/client"
	"github.com/t0k4rt/livebox-go/client/ip_address"
)

type liveBoxProvider struct {
	liveboxClient *liveboxApi.Livebox
}

func NewIPProvider() *liveBoxProvider {
	return &liveBoxProvider{
		liveboxClient: liveboxApi.Default,
	}
}

func (l *liveBoxProvider) GetIP() (*ipprovider.ProvidedIP, error) {

	resp, err := l.liveboxClient.IPAddress.PostSysbusNMCGetWANStatus(ip_address.NewPostSysbusNMCGetWANStatusParams())
	if err != nil {
		return nil, err
	}

	return ipprovider.NewFromString(
		resp.Payload.Result.Data.IPAddress,
		resp.Payload.Result.Data.IPV6Address,
	), nil
}
