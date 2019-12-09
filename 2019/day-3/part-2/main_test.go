package main

import (
	"reflect"
	"testing"
)

func Test_genPath(t *testing.T) {
	type args struct {
		pathString []string
	}
	tests := []struct {
		name string
		args args
		// NOTE: It fails if I change this to slice of pointers
		want []Coordinate
	}{
		// TODO: Add test cases.
		{
			"U7,R6,D4,L4",
			args{
				pathString: []string{"U7", "R6", "D4", "L4"},
			},
			[]Coordinate{
				Coordinate{0, 0},
				Coordinate{0, 7},
				Coordinate{6, 7},
				Coordinate{6, 3},
				Coordinate{2, 3},
			},
		},
		{
			"R8,U5,L5,D3",
			args{
				pathString: []string{"R8", "U5", "L5", "D3"},
			},
			[]Coordinate{
				Coordinate{0, 0},
				Coordinate{8, 0},
				Coordinate{8, 5},
				Coordinate{3, 5},
				Coordinate{3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genPath(tt.args.pathString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genPath() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_getDirectionAndDistance(t *testing.T) {
	type args struct {
		pathString string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		// TODO: Add test cases.
		{
			"U7",
			args{
				"U7",
			},
			"U",
			7,
		},
		{
			"R6",
			args{
				"R6",
			},
			"R",
			6,
		},
		{
			"D4",
			args{
				"D4",
			},
			"D",
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getDirectionAndDistance(tt.args.pathString)
			if got != tt.want {
				t.Errorf("getDirectionAndDistance() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getDirectionAndDistance() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findIntersection(t *testing.T) {
	type args struct {
		path1 [2]Coordinate
		path2 [2]Coordinate
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 Coordinate
	}{
		// TODO: Add test cases.
		{
			"0,0-0,3 and -1,1-1,1",
			args{
				[2]Coordinate{
					Coordinate{0, 0},
					Coordinate{0, 3},
				},
				[2]Coordinate{
					Coordinate{-1, 1},
					Coordinate{1, 1},
				},
			},
			true,
			Coordinate{
				0, 1,
			},
		},
		{
			"0,0-0,7 and 0,0-8,0",
			args{
				[2]Coordinate{
					Coordinate{0, 0},
					Coordinate{0, 7},
				},
				[2]Coordinate{
					Coordinate{0, 0},
					Coordinate{8, 0},
				},
			},
			false,
			Coordinate{},
		},
		{
			"0,7-6,7 and 0,0-8,0",
			args{
				[2]Coordinate{
					Coordinate{0, 7},
					Coordinate{6, 7},
				},
				[2]Coordinate{
					Coordinate{0, 0},
					Coordinate{8, 0},
				},
			},
			false,
			Coordinate{},
		},
		{
			"0,0-0,3 and 2,4-4,4",
			args{
				[2]Coordinate{
					Coordinate{0, 0},
					Coordinate{0, 3},
				},
				[2]Coordinate{
					Coordinate{2, 4},
					Coordinate{4, 4},
				},
			},
			false,
			Coordinate{},
		},
		{
			"3,5-3,2 and 6,3-2,3",
			args{
				[2]Coordinate{
					Coordinate{3, 5},
					Coordinate{3, 2},
				},
				[2]Coordinate{
					Coordinate{6, 3},
					Coordinate{2, 3},
				},
			},
			true,
			Coordinate{3, 3},
		},
		{
			"0,7-6,7 and 8,0-8,5",
			args{
				[2]Coordinate{
					Coordinate{0, 7},
					Coordinate{6, 7},
				},
				[2]Coordinate{
					Coordinate{8, 0},
					Coordinate{8, 5},
				},
			},
			false,
			Coordinate{},
		},
		{
			"6,7-6,3 and 8,0-8,5",
			args{
				[2]Coordinate{
					Coordinate{6, 7},
					Coordinate{6, 3},
				},
				[2]Coordinate{
					Coordinate{8, 0},
					Coordinate{8, 5},
				},
			},
			false,
			Coordinate{},
		},
		{
			"6,3-2,3 and 8,0-8,5",
			args{
				[2]Coordinate{
					Coordinate{6, 3},
					Coordinate{2, 3},
				},
				[2]Coordinate{
					Coordinate{8, 0},
					Coordinate{8, 5},
				},
			},
			false,
			Coordinate{},
		},
		{
			"0,0-8,0 and 6,7-6,3",
			args{
				[2]Coordinate{
					Coordinate{0, 0},
					Coordinate{8, 0},
				},
				[2]Coordinate{
					Coordinate{6, 7},
					Coordinate{6, 3},
				},
			},
			false,
			Coordinate{},
		},
		{
			"3,5-3,2 and 0,7-67",
			args{
				[2]Coordinate{
					Coordinate{3, 5},
					Coordinate{3, 2},
				},
				[2]Coordinate{
					Coordinate{0, 7},
					Coordinate{6, 7},
				},
			},
			false,
			Coordinate{},
		},
		{
			"6,7-6,3 and 8,5-3,5",
			args{
				[2]Coordinate{
					Coordinate{6, 7},
					Coordinate{6, 3},
				},
				[2]Coordinate{
					Coordinate{8, 5},
					Coordinate{3, 5},
				},
			},
			true,
			Coordinate{6, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findIntersection(tt.args.path1, tt.args.path2)
			if got != tt.want {
				t.Errorf("findIntersection() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("findIntersection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findSteps(t *testing.T) {
	type args struct {
		steps []string
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"simple",
			args{
				[]string{"R8", "U5", "L5", "D3"},
				3,
			},
			18,
		},
		{
			"simple2",
			args{
				[]string{"U7", "R6", "D4", "L4"},
				3,
			},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSteps(tt.args.steps, tt.args.index); got != tt.want {
				t.Errorf("findSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findExtra(t *testing.T) {
	type args struct {
		intersection Coordinate
		index        int
		wire         Wire
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			"simple",
			args{
				Coordinate{
					6,
					5,
				},
				2,
				Wire{
					Coordinates: []Coordinate{
						Coordinate{0, 0},
						Coordinate{8, 0},
						Coordinate{8, 5},
						Coordinate{3, 5},
						Coordinate{3, 2},
					},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findExtra(tt.args.intersection, tt.args.index, tt.args.wire); got != tt.want {
				t.Errorf("findExtra() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_findExtra(t *testing.T) {
// 	type args struct {
// 		intersection   Coordinate
// 		lastCoordinate Coordinate
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want float64
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			"simple",
// 			args{
// 				Coordinate{
// 					6,
// 					5,
// 				},
// 				Coordinate{
// 					8,
// 					5,
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := findExtra(tt.args.intersection, tt.args.lastCoordinate); got != tt.want {
// 				t.Errorf("findExtra() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
