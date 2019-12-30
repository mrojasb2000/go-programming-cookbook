package collections

import (
	"reflect"
	"testing"
)

func TestLowerCaseData(t *testing.T) {
	type args struct {
		w WorkWith
	}
	tests := []struct {
		name string
		args args
		want WorkWith
	}{
		{"base-case", args{WorkWith{"Test", 0}}, WorkWith{"test", 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCaseData(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LowerCaseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
