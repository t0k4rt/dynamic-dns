package ipv6

import (
	"net"
	"reflect"
	"testing"
)

func TestIsPublicIPv6(t *testing.T) {
	type args struct {
		ipString string
	}
	tests := []struct {
		name    string
		args    args
		want    net.IP
		wantErr bool
	}{
		{"test ipv4", args{"192.168.1.2"}, nil, true},
		{"test ipv4 cidr", args{"192.168.1.2/16"}, nil, true},
		{"test ipv6", args{"2a01:cb00:c30:700:c9e:cc9b:ba33:3022"}, net.ParseIP("2a01:cb00:c30:700:c9e:cc9b:ba33:3022"), false},
		{"test ipv6 cidr", args{"2a01:cb00:c30:700:c9e:cc9b:ba33:3060/126"}, net.ParseIP("2a01:cb00:c30:700:c9e:cc9b:ba33:3060"), false},
		{"test ipv6 private", args{"fe80::10a8:478d:3e51:6438"}, nil, true},
		{"test ipv6 private cidr", args{"fe80::10a8:478d:3e51:6438/64"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePublicIPv6(tt.args.ipString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePublicIPv6() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePublicIPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}
