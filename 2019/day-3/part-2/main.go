package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
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

type Intersection struct {
	Coordinate
	Distance int
}

func main() {
	start := time.Now()
	answer := findAnswer()
	fmt.Printf("The answer is: %+v. That took %s\n", int(answer), time.Since(start))
}

func findAnswer() int {
	wire1Coordinates := inputs[0].Coordinates
	wire1Paths := inputs[0].Paths
	wire2Coordinates := inputs[1].Coordinates
	wire2Paths := inputs[1].Paths

	var intersections []Intersection
	var minDistance int
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
				wire1Steps := findSteps(wire1Paths, i)
				wire1Steps = wire1Steps + int(findExtra(intersection, inputs[0].Coordinates[i]))
				wire2Steps := findSteps(wire2Paths, j)
				wire2Steps = wire2Steps + int(findExtra(intersection, inputs[1].Coordinates[j]))
				totalDistance := wire1Steps + wire2Steps

				// Set the first intersection's distance as minDistance
				if minDistance == 0 {
					minDistance = totalDistance
				} else {
					if minDistance > totalDistance {
						minDistance = totalDistance
					}
				}

				intersections = append(intersections, Intersection{
					intersection,
					wire1Steps + wire2Steps,
				})
			}
		}
	}

	return minDistance
}

func findExtra(intersection, lastCoordinate Coordinate) float64 {
	return math.Abs(float64(intersection.X-lastCoordinate.X)) + math.Abs(float64(intersection.Y-lastCoordinate.Y))
}

func findSteps(steps []string, index int) int {
	var totalDistance int
	for i := 0; i < index; i++ {
		_, distance := getDirectionAndDistance(steps[i])
		totalDistance = totalDistance + distance
	}
	return totalDistance
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
