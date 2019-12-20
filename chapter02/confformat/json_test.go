package confformat

import (
	"testing"
	"reflect"
    "bytes"
)

func TestJSONData_ToJSON(t *testing.T) {
	type fields struct {
		Name string
		Age int
	}
	tests := []struct {
		name string
		fields fields
		want *bytes.Buffer
		wantErr bool
	}{
		{"base-case", fields{"Someone", 100}, bytes.NewBufferString(`{"name":"Someone","age":100}`), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			td := &JSONData{
				Name: tt.fields.Name,
				Age: tt.fields.Age,
			}
			got, err := td.ToJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONData.ToJSON() err = %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONData.ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
