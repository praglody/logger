package logger

import "testing"

func TestInit(t *testing.T) {
	type args struct {
		logfile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"init", args{
			logfile: "./logs/main.log",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.logfile); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				Info("sagasfasdfsd")
			}
		})
	}
}
