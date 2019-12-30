package tags

import "testing"

func TestEmptyStruct(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EmptyStruct(); (err != nil) != tt.wantErr {
				t.Errorf("EmptyStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
