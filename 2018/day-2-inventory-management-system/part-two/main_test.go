package main

import (
	"testing"
)

func Test_charDiff(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"abcde,axcye",
			args{
				"abcde",
				"axcye",
			},
			2,
		},
		{
			"fghij,fguij",
			args{
				"fghij",
				"fguij",
			},
			1,
		},
		{
			"pqrst,fguij",
			args{
				"pqrst",
				"fguij",
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charDiff(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("charDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"abcde,axcye",
			args{
				"abcde",
				"axcye",
			},
			"ace",
		},
		{
			"fghij,fguij",
			args{
				"fghij",
				"fguij",
			},
			"fgij",
		},
		{
			"pqrst,fguij",
			args{
				"pqrst",
				"fguij",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAnswer(t *testing.T) {
	type args struct {
		sortedInputArray []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"abcde fghij klmno pqrst fguij axcye wvxyz",
			args{
				[]string{
					"abcde",
					"axcye",
					"fghij",
					"fguij",
					"klmno",
					"pqrst",
					"wvxyz",
				},
			},
			"fgij",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnswer(tt.args.sortedInputArray); got != tt.want {
				t.Errorf("findAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}
