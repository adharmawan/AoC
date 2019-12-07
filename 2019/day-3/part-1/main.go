package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var inputs []Wire

type Wire struct {
	Paths       []string
	Coordinates []Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func main() {
	answer := findAnswer()
	fmt.Printf("The answer is: %+v\n", int(answer))
}

func findAnswer() float64 {
	wire1Coordinates := inputs[0].Coordinates
	wire2Coordinates := inputs[1].Coordinates

	var intersections []Coordinate
	for i := 0; i < len(wire1Coordinates)-1; i++ {
		for j := 0; j < len(wire2Coordinates)-1; j++ {
			found, intersection := findIntersection([2]Coordinate{
				wire1Coordinates[i],
				wire1Coordinates[i+1],
			}, [2]Coordinate{
				wire2Coordinates[j],
				wire2Coordinates[j+1],
			})

			if found {
				intersections = append(intersections, intersection)
			}
		}
	}

	var distances []float64
	for _, intersection := range intersections {
		distance := fromCoordinatesToDistance(intersection)
		distances = append(distances, distance)
	}

	var minDistance float64
	minDistance = distances[0]
	for _, distance := range distances {
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func findIntersections() {

}
func fromCoordinatesToDistance(coordinate Coordinate) float64 {
	return math.Abs(float64(coordinate.X)) + math.Abs(float64(coordinate.Y))
}

func findIntersection(path1, path2 [2]Coordinate) (bool, Coordinate) {
	toCompareFirst := path1
	toCompareSecond := path2

	// Find the X that does not change
	if path1[0].X-path1[1].X != 0 {
		toCompareFirst = path2
		toCompareSecond = path1
	}

	// Check if it's between the Xs on the other wire
	smallerX := toCompareSecond[0]
	biggerX := toCompareSecond[1]
	smallerY := toCompareFirst[0]
	biggerY := toCompareFirst[1]

	if toCompareSecond[0].X > toCompareSecond[1].X {
		smallerX = toCompareSecond[1]
		biggerX = toCompareSecond[0]
	}

	if toCompareFirst[0].Y > toCompareFirst[1].Y {
		smallerY = toCompareFirst[1]
		biggerY = toCompareFirst[0]
	}
	if toCompareFirst[0].X > smallerX.X && toCompareFirst[0].X < biggerX.X && toCompareSecond[0].Y > smallerY.Y && toCompareSecond[0].Y < biggerY.Y {
		return true, Coordinate{toCompareFirst[0].X, toCompareSecond[0].Y}
	}

	return false, Coordinate{}
}



// From U11 to {"U", 11}
func getDirectionAndDistance(pathString string) (string, int) {
	direction := string(pathString[0])
	distance, err := strconv.Atoi(pathString[1:])
	if err != nil {
		panic(err)
	}
	return direction, distance
}

// From U11, R4, D3 to [{0,0},{0,11},{4,11},{4,8}]
func genPath(pathStrings []string) []Coordinate {
	var coordinates []Coordinate
	coordinate := Coordinate{0, 0}
	for _, pathString := range pathStrings {
		direction, distance := getDirectionAndDistance(pathString)
		if direction == "R" {
			coordinate.X = coordinate.X + distance
		} else if direction == "L" {
			coordinate.X = coordinate.X - distance
		} else if direction == "U" {
			coordinate.Y = coordinate.Y + distance
		} else if direction == "D" {
			coordinate.Y = coordinate.Y - distance
		}
		coordinates = append(coordinates, coordinate)
	}
	return append([]Coordinate{Coordinate{0, 0}}, coordinates...)
}

func init() {
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	wires := strings.FieldsFunc(string(inputByte), func(r rune) bool {
		return r == '\n'
	})

	for _, wire := range wires {
		paths := strings.FieldsFunc(string(wire), func(r rune) bool {
			return r == ','
		})

		coordinates := genPath(paths)

		inputs = append(inputs, Wire{
			paths,
			coordinates,
		})
	}
}
