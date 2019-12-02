package main

import "testing"

func Test_fuelNeeded(t *testing.T) {
	type args struct {
		mass float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"12",
			args{
				12,
			},
			2,
		},
		{
			"14",
			args{
				14,
			},
			2,
		},
		{
			"1969",
			args{
				1969,
			},
			654,
		},
		{
			"100756",
			args{
				100756,
			},
			33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuelNeeded(tt.args.mass); got != tt.want {
				t.Errorf("fuelNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}
