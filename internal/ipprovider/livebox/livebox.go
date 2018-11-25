package livebox

import (
	"net"

	liveboxApi "github.com/t0k4rt/livebox-go/client"
	"github.com/t0k4rt/livebox-go/client/ip_address"
)

type liveBoxProvider struct {
	liveboxClient *liveboxApi.Livebox
}

func NewIPProvider() (*liveBoxProvider, error) {
	return &liveBoxProvider{
		liveboxClient: liveboxApi.Default,
	}, nil
}

// func (l *liveBoxProvider) GetIP() (*ipprovider.ProvidedIP, error) {

// 	return ipprovider.NewFromString(
// 		resp.Payload.Result.Data.IPAddress,
// 		resp.Payload.Result.Data.IPV6Address,
// 	), nil
// }

func (l *liveBoxProvider) GetIP(version int) (net.IP, error) {
	var err error
	var ip net.IP
	switch version {
	case 4:
		ip, err = l.GetIPv4()
	case 6:
		ip, err = l.GetIPv6()
	}
	return ip, err
}

func (l *liveBoxProvider) GetIPv4() (net.IP, error) {

	resp, err := l.liveboxClient.IPAddress.PostSysbusNMCGetWANStatus(ip_address.NewPostSysbusNMCGetWANStatusParams())
	if err != nil {
		return nil, err
	}
	return net.ParseIP(resp.Payload.Result.Data.IPAddress), nil

}

func (l *liveBoxProvider) GetIPv6() (net.IP, error) {

	resp, err := l.liveboxClient.IPAddress.PostSysbusNMCGetWANStatus(ip_address.NewPostSysbusNMCGetWANStatusParams())
	if err != nil {
		return nil, err
	}
	return net.ParseIP(resp.Payload.Result.Data.IPV6Address), nil
}
