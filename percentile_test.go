package gf

import "testing"

type T struct {
	name   string
	number int
}

func TestPercentile(t *testing.T) {
	type args struct {
		percentile int
		slice      []T
		less       func(slice []T, i, j int) bool
		get        func(slice []T, i int) float64
		avg        func(slice []T, i, j int) float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				percentile: 50,
				slice: []T{
					{
						name:   "a",
						number: 1,
					},
					{
						name:   "b",
						number: 10,
					},
					{
						name:   "f",
						number: 1000,
					},
					{
						name:   "c",
						number: 9,
					},
					{
						name:   "d",
						number: 123,
					},
					{
						name:   "e",
						number: 444,
					},
				},
				less: func(slice []T, i, j int) bool {
					return slice[i].number < slice[j].number
				},
				get: func(slice []T, i int) float64 {
					return float64(slice[i].number)
				},
				avg: func(slice []T, i, j int) float64 {
					return float64(slice[i].number+slice[j].number) / 2
				},
			},
			want:    66.5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Percentile(tt.args.percentile, tt.args.slice, tt.args.less, tt.args.get, tt.args.avg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Percentile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Percentile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
