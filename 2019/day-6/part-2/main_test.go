package main

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	type args struct {
		left  []string
		right []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"simple",
			args{
				[]string{"a", "b"},
				[]string{"b", "c"},
			},
			[]string{"a", "b", "c"},
		},
		{
			"simple_reverse",
			args{
				[]string{"b", "c"},
				[]string{"a", "b"},
			},
			[]string{"b", "c", "a"},
		},
		{
			"inputeasy",
			args{
				[]string{"COM", "B", "C", "D", "E", "B", "G", "D", "E", "J", "K"},
				[]string{"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"},
			},
			[]string{"COM", "B", "C", "D", "E", "G", "J", "K", "F", "H", "I", "L"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIndex(t *testing.T) {
	type args struct {
		array []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple",
			args{
				[]string{"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"},
				"B",
			},
			0,
		},
		{
			"simple",
			args{
				[]string{"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"},
				"C",
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndex(tt.args.array, tt.args.item); got != tt.want {
				t.Errorf("findIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePathToCOM(t *testing.T) {
	type args struct {
		left  []string
		right []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"simple",
			args{
				[]string{"COM", "B", "C", "D", "E", "B", "G", "D", "E", "J", "K", "K", "I"},
				[]string{"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "YOU", "SAN"},
				"YOU",
			},
			[]string{"COM", "B", "C", "D", "E", "J", "K", "YOU"},
		},
		{
			"simple",
			args{
				[]string{"COM", "B", "C", "D", "E", "B", "G", "D", "E", "J", "K", "K", "I"},
				[]string{"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "YOU", "SAN"},
				"SAN",
			},
			[]string{"COM", "B", "C", "D", "I", "SAN"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePathToCOM(tt.args.left, tt.args.right, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePathToCOM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIntersectionPoint(t *testing.T) {
	type args struct {
		you []string
		san []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple",
			args{
				[]string{"COM", "B", "C", "D", "E", "J", "K", "YOU"},
				[]string{"COM", "B", "C", "D", "I", "SAN"},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntersectionPoint(tt.args.you, tt.args.san); got != tt.want {
				t.Errorf("findIntersectionPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
