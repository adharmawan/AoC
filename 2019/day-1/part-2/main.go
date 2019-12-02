package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var (
	inputs []float64
)

func main() {
	answer := findAnswer(inputs)
	fmt.Printf("The answer is: %+v\n", int(answer))
}

func findAnswer(inputs []float64) float64 {
	var total float64
	for _, input := range inputs {
		fuel := fuelNeeded(input)
		total = total + fuel + fuelNeededByFuel(fuel)
	}
	return total
}

func fuelNeededByFuel(fuel float64) float64 {
	total := fuelNeeded(fuel)
	if total <= 0 {
		return 0
	}
	return total + fuelNeededByFuel(total)
}

func init() {
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputArray := strings.Fields(string(inputByte))
	inputs, err = parseStringArrayToFloat64(inputArray)
	if err != nil {
		panic(err)
	}
}

func parseStringArrayToFloat64(stringArray []string) ([]float64, error) {
	float64Array := make([]float64, len(stringArray))
	var err error
	for i := 0; i < len(stringArray); i++ {
		float64Array[i], err = strconv.ParseFloat(stringArray[i], 10)
		if err != nil {
			return nil, err
		}
	}
	return float64Array, nil
}

func fuelNeeded(mass float64) float64 {
	return math.Floor(mass/3) - 2
}
