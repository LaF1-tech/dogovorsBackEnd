package env

import (
	"reflect"
	"testing"
)

func Test_parseEnvFileData(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "valid",
			args: args{
				content: "ENV1=\"aaaaa\"\nENV2='bbbbb'\nENV3=ccccc\n",
			},
			want: map[string]string{"ENV1": "aaaaa", "ENV2": "bbbbb", "ENV3": "ccccc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseEnvData(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEnvData() = %v, want %v", got, tt.want)
			}
		})
	}
}
