package log

import "testing"

func TestOrifinalError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OriginalError(); (err != nil) != tt.wantErr {
				t.Errorf("OriginalError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestThroughError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PassThroughError(); (err != nil) != tt.wantErr {
				t.Errorf("PassThroughError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
