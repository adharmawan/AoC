package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"1,0,0,0,99",
			args{
				[]int{1, 0, 0, 0, 99},
			},
			[]int{2, 0, 0, 0, 99},
			false,
		},
		{
			"2,3,0,3,99",
			args{
				[]int{2, 3, 0, 3, 99},
			},
			[]int{2, 3, 0, 6, 99},
			false,
		},
		{
			"2,4,4,5,99,0",
			args{
				[]int{2, 4, 4, 5, 99, 0},
			},
			[]int{2, 4, 4, 5, 99, 9801},
			false,
		},
		{
			"1,1,1,4,99,5,6,0,99",
			args{
				[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.args.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
