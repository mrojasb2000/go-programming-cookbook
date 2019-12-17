package envvar

import "testing"

func TestLoadConfig(t *testing.T) {
	c := struct {
		Example string `json:"example"`
	}{}
	type args struct {
		path      string
		envPrefix string
		config    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"123", "", &c}, true},
		{"no-file", args{"", "", &c}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadConfig(tt.args.path, tt.args.envPrefix, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
