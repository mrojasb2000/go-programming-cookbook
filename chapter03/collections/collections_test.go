package collections

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	ws := []WorkWith{WorkWith{"test", 1}}

	type args struct {
		ws []WorkWith
		f  func(w WorkWith) bool
	}
	tests := []struct {
		name string
		args args
		want []WorkWith
	}{
		{"filter true", args{ws, func(w WorkWith) bool { return true }}, []WorkWith{WorkWith{"test", 1}}},
		{"filter false", args{ws, func(w WorkWith) bool { return false }}, []WorkWith{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.ws, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
