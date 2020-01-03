package structured

import "testing"

func TestThrowError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ThrowError(); (err != nil) != tt.wantErr {
				t.Errorf("ThrowError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
