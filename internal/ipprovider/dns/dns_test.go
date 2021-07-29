package dns

import (
	"fmt"
	"testing"
)

func TestGetIPv4(t *testing.T) {
	dnsprovider, err := NewIPProvider()
	if err != nil {
		fmt.Println(err)
	}
	ip, err := dnsprovider.GetIPv4()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip.String())
}
func TestGetIPv6(t *testing.T) {
	dnsprovider, err := NewIPProvider()
	if err != nil {
		fmt.Println(err)
	}
	ip, err := dnsprovider.GetIPv6()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip.String())
}
