package gf

import "testing"

func TestFloat32KeepDecimal(t *testing.T) {
	type args struct {
		floatValue   float32
		decimalCount int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success",
			args: args{
				floatValue:   123.456789,
				decimalCount: 2,
			},
			want: 123.46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatKeepDecimal(tt.args.floatValue, tt.args.decimalCount); got != tt.want {
				t.Errorf("FloatKeepDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64KeepDecimal(t *testing.T) {
	type args struct {
		floatValue   float64
		decimalCount int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success",
			args: args{
				floatValue:   123.456789,
				decimalCount: 2,
			},
			want: 123.46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatKeepDecimal(tt.args.floatValue, tt.args.decimalCount); got != tt.want {
				t.Errorf("FloatKeepDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
