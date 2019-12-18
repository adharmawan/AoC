package main

import (
	"reflect"
	"testing"
)

func Test_twoAdjacentDigitNotPartOfLargerGroup(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"112233",
			args{
				"112233",
			},
			true,
		},
		{
			"123444",
			args{
				"123444",
			},
			false,
		},
		{
			"111122",
			args{
				"111122",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoAdjacentDigitNotPartOfLargerGroup(tt.args.num); got != tt.want {
				t.Errorf("twoAdjacentDigitNotPartOfLargerGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findUnique(t *testing.T) {
	type args struct {
		array []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"[49,50,51]",
			args{
				[]byte{49, 50, 51},
			},
			[]byte{49, 50, 51},
		},
		{
			"[52,52]",
			args{
				[]byte{52, 52},
			},
			[]byte{},
		},
		{
			"[49,49,49,50]",
			args{
				[]byte{49, 49, 49, 50},
			},
			[]byte{50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnique(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
