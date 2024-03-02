package randomizer

import "testing"

func TestRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{length: 10},
			want:    10,
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{length: 100},
			want:    100,
			wantErr: false,
		},
		{
			name:    "test3",
			args:    args{length: 1000},
			want:    1000,
			wantErr: false,
		},
		{
			name:    "test4",
			args:    args{length: 10000},
			want:    10000,
			wantErr: false,
		},
		{
			name:    "test5",
			args:    args{length: -1},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomString(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("RandomString() len of got = %v, want %v", got, tt.want)
			}

			t.Logf("RandomString() got = %v, len = %v", got, len(got))
		})
	}
}
