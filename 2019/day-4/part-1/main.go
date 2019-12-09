package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func isSixDigit(password int) bool {
	passwordString := strconv.Itoa(password)
	match, err := regexp.MatchString(`^\d{6}$`, passwordString)
	if err != nil {
		return false
	}
	return match
}

func isWithinRange(password int, min, max int) bool {
	// return password > min && password < min
	return password >= min && password <= max
}

func twoAdjacentDigitsAreTheSame(password int) bool {
	passwordString := strconv.Itoa(password)

	first := int(passwordString[0])
	second := int(passwordString[1])
	third := int(passwordString[2])
	fourth := int(passwordString[3])
	fifth := int(passwordString[4])
	sixth := int(passwordString[5])
	return first == second || second == third || third == fourth || fourth == fifth || fifth == sixth
}

func digitsNeverDecrease(password int) bool {
	passwordString := strconv.Itoa(password)

	first := int(passwordString[0])
	second := int(passwordString[1])
	third := int(passwordString[2])
	fourth := int(passwordString[3])
	fifth := int(passwordString[4])
	sixth := int(passwordString[5])
	return first <= second && second <= third && third <= fourth && fourth <= fifth && fifth <= sixth
}

func valid(password, min, max int) bool {
	return isSixDigit(password) &&
		isWithinRange(password, min, max) &&
		twoAdjacentDigitsAreTheSame(password) &&
		digitsNeverDecrease(password)
}

func findAnswer() int {
	var total int
	for i := input.Min; i <= input.Max; i++ {
		if valid(i, input.Min, input.Max) {
			total = total + 1
		}
	}
	return total
}
func main() {
	fmt.Println(input)
	answer := findAnswer()
	fmt.Printf("The answer is: %+v\n", int(answer))
}

type Input struct {
	Min int
	Max int
}

var input Input

func init() {
	var err error
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rangeOfNumbers := strings.FieldsFunc(string(inputByte), func(r rune) bool {
		return r == '-'
	})

	input.Min, err = strconv.Atoi(rangeOfNumbers[0])
	if err != nil {
		panic(err)
	}

	input.Max, err = strconv.Atoi(rangeOfNumbers[1])
	if err != nil {
		panic(err)
	}
}
