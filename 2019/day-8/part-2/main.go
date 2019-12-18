package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func findAnswer() string {
	area := width * height
	var picture string
	for i := 0; i < area; i++ {
		if i%width == 0 {
			picture += "\n"
		}
		picture += findColor(i, area, input)
	}
	return picture
}

func findColor(index, area int, str string) string {
	char := string(str[index])
	if char != "2" {
		if char == "0" {
			return " "
		}
		return char
	}
	return findColor(index, area, str[area:])
}
func main() {
	start := time.Now()
	answer := findAnswer()
	fmt.Printf("The answer is: \n%+v\n\nThat took %s\n", answer, time.Since(start))
}

var input string
var width = 25
var height = 6

func init() {
	var err error
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input = string(inputByte)
}
