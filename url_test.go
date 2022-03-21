package gf

import "testing"

func TestUrlFormat(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success-1",
			args:    args{urlStr: "https://u:p@www.a.com:9999////////b/c/d.e?i=2&&i=1&&&&&&&&g=2&h=3&h=3&h=3&f=4#5"},
			want:    "https://u:p@www.a.com:9999/b/c/d.e?f=4&g=2&h=3&i=1&i=2#5",
			wantErr: false,
		},
		{
			name:    "success-2",
			args:    args{urlStr: "https://u:p@www.a.com:9999/#/b/c/d.e?i=2&&i=1&&&&&&&&g=2&h=3&h=3&h=3&f=4"},
			want:    "https://u:p@www.a.com:9999/#/b/c/d.e?i=2&&i=1&&&&&&&&g=2&h=3&h=3&h=3&f=4",
			wantErr: false,
		},
		{
			name:    "failed-1",
			args:    args{urlStr: "https://www.a .com/b/c/d"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "failed-2",
			args:    args{urlStr: "https:///b/c/d"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UrlFormat(tt.args.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("UrlFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UrlFormat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
