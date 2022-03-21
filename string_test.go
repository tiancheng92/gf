package gf

import (
	"reflect"
	"testing"
)

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				b: []byte("test"),
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "success",
			args: args{
				s: "test",
			},
			want: []byte("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringCreateGzip(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				str: "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			},
			want:    []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 50, 52, 50, 54, 49, 53, 51, 183, 176, 52, 160, 14, 11, 0, 0, 0, 255, 255, 1, 0, 0, 255, 255, 114, 74, 169, 124, 80, 0, 0, 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringCreateGzip(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringCreateGzip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringCreateGzip() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringParseGzip(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				data: []byte{31, 139, 8, 0, 0, 0, 0, 0, 0, 255, 50, 52, 50, 54, 49, 53, 51, 183, 176, 52, 160, 14, 11, 0, 0, 0, 255, 255, 1, 0, 0, 255, 255, 114, 74, 169, 124, 80, 0, 0, 0},
			},
			want:    "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringParseGzip(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringParseGzip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringParseGzip() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringJoin(t *testing.T) {
	type args struct {
		variables []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				[]string{"a", "b"},
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringJoin(tt.args.variables...); got != tt.want {
				t.Errorf("StringJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
