package main

import (
	"testing"
)

func Test_hasTwos(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcdef",
			args{
				id: "abcdef",
			},
			false,
		},
		{
			"bababc",
			args{
				id: "bababc",
			},
			true,
		},
		{
			"abbcde",
			args{
				id: "abbcde",
			},
			true,
		},
		{
			"abcccd",
			args{
				id: "abcccd",
			},
			false,
		},
		{
			"aabcdd",
			args{
				id: "aabcdd",
			},
			true,
		},
		{
			"abcdee",
			args{
				id: "abcdee",
			},
			true,
		},
		{
			"ababab",
			args{
				id: "ababab",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasTwos(tt.args.id); got != tt.want {
				t.Errorf("hasTwos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasThrees(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcdef",
			args{
				id: "abcdef",
			},
			false,
		},
		{
			"bababc",
			args{
				id: "bababc",
			},
			true,
		},
		{
			"abbcde",
			args{
				id: "abbcde",
			},
			false,
		},
		{
			"abcccd",
			args{
				id: "abcccd",
			},
			true,
		},
		{
			"aabcdd",
			args{
				id: "aabcdd",
			},
			false,
		},
		{
			"abcdee",
			args{
				id: "abcdee",
			},
			false,
		},
		{
			"ababab",
			args{
				id: "ababab",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasThrees(tt.args.id); got != tt.want {
				t.Errorf("hasThrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
