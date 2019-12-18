package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func findAnswer() int {
	width := 25
	height := 6
	toSkip := width * height
	answer := Layer{}
	for i := 0; i <= len(input); i = i + toSkip {
		end := i + toSkip
		if len(input) < end {
			end = len(input)
		}
		layerString := input[i:end]
		layerWithoutZeroes := removeZeroes(layerString)
		layer := toLayer(layerWithoutZeroes)

		if answer.LengthWithoutZeroes < layer.LengthWithoutZeroes {
			answer = layer
		}

	}
	return answer.SumOfOneDigitsAndTwoDigitsWithoutZeroes
}

type Layer struct {
	SumOfOneDigitsAndTwoDigitsWithoutZeroes int
	LengthWithoutZeroes                     int
}

func removeZeroes(pixel string) string {
	newLayer := ""
	for j := 0; j < len(pixel); j++ {
		if pixel[j] != '0' {
			newLayer = newLayer + string(pixel[j])
		}
	}
	return newLayer
}

func toLayer(pixelWithoutZeroes string) Layer {
	var numOfOnes int
	var numOfTwos int
	for j := 0; j < len(pixelWithoutZeroes); j++ {
		temp, _ := strconv.Atoi(string(pixelWithoutZeroes[j]))
		if temp == 1 {
			numOfOnes++
		} else if temp == 2 {
			numOfTwos++
		}
	}

	return Layer{
		SumOfOneDigitsAndTwoDigitsWithoutZeroes: numOfOnes * numOfTwos,
		LengthWithoutZeroes:                     len(pixelWithoutZeroes),
	}
}

func main() {
	start := time.Now()
	answer := findAnswer()
	fmt.Printf("The answer is: %+v. That took %s\n", int(answer), time.Since(start))
}

var input string

func init() {
	var err error
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input = string(inputByte)
}
