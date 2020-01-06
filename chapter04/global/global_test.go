package global

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetLog(t *testing.T) {
	type args struct {
		l *logrus.Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{logrus.New()}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLog(tt.args.l)
		})
	}
}
