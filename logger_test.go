package logger

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	type args struct {
		logname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "=====",
			args: args{
				logname: "data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.logname); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				Infof("%s\n", time.Now().Format(time.DateTime))
			}
		})
	}
}
