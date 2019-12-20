package confformat

import (
	"bytes"
	"reflect"
	"testing"
)

func TestTOMLData_ToTOML(t *testing.T) {
	type fields struct {
		Name string
		Age  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *bytes.Buffer
		wantErr bool
	}{
		{"base-case", fields{"Someone", 100}, bytes.NewBufferString("name = \"Someone\"\nage = 100\n"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := &TOMLData{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}

			got, err := td.ToTOML()
			if (err != nil) != tt.wantErr {
				t.Errorf("TOMLData.ToTOML() error = %v, want %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("TOMLData.ToTOML() = %v, want %v", got, tt.want)
			}
		})
	}
}
