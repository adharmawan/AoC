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
