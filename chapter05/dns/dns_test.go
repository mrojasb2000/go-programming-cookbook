package dns

import "testing"

func TestLookupAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"www.google.com"}, false},
		{"case 2", args{"A"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LookupAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookupAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
