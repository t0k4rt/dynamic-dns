package domainparser

import "testing"

func TestParseDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{"test toktok.fr", args{"toktok.fr"}, "toktok.fr", "@", false},
		{"test www.toktok.fr", args{"www.toktok.fr"}, "toktok.fr", "www", false},
		{"test www.plop.toktok.fr", args{"www.plop.toktok.fr"}, "toktok.fr", "www.plop", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseDomain(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDomain() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseDomain() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
