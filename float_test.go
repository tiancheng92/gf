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
			if got := FloatRound(tt.args.floatValue, tt.args.decimalCount); got != tt.want {
				t.Errorf("FloatRound() = %v, want %v", got, tt.want)
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
			if got := FloatRound(tt.args.floatValue, tt.args.decimalCount); got != tt.want {
				t.Errorf("FloatRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeFloat64KeepDecimal(t *testing.T) {
	type typeFloat64 float64
	type args struct {
		floatValue   typeFloat64
		decimalCount int
	}
	tests := []struct {
		name string
		args args
		want typeFloat64
	}{
		{
			name: "success",
			args: args{
				floatValue:   typeFloat64(123.456789),
				decimalCount: 2,
			},
			want: typeFloat64(123.46),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatRound(tt.args.floatValue, tt.args.decimalCount); got != tt.want {
				t.Errorf("FloatRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
