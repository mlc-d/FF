package auth

import (
	"testing"
)

func TestBuildToken(t *testing.T) {
	var a int64 = 2
	var b uint8 = 3
	type args struct {
		userID   *int64
		userRole *uint8
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "testubi1",
			args:    args{userID: &a, userRole: &b},
			want:    "nada",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildToken(tt.args.userID, tt.args.userRole)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
