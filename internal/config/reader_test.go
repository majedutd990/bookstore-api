package config

import "testing"

func testParse(t *testing.T) {
	type args struct {
		path string
		cfg  *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "parse config.yaml",
			args: args{
				path: "testdata/config.yaml",
				cfg:  &Config{},
			},
			wantErr: false,
		},
		{
			name: "parse config.yml",
			args: args{
				path: "testdata/config.yml",
				cfg:  &Config{},
			},
			wantErr: false,
		},
		{
			name: "error file extension not allowed",
			args: args{
				path: "testdata/config.conf",
				cfg:  &Config{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Parser(tt.args.path, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Parser() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
