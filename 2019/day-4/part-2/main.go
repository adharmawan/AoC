package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	min int
	max int
)

func main() {
	answer := findAnswer(min, max)
	fmt.Printf("The answer is: %+v\n", int(answer))
}

func findAnswer(min, max int) int {
	found := 0
	for i := min; i < max; i++ {
		if isValid(i, min, max) {
			found = found + 1
		}
	}
	return found
}

func twoAdjacentDigit(num string) bool {
	found := false
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			found = true
			break
		}
	}
	return found
}

func twoAdjacentDigitNotPartOfLargerGroup(num string) bool {
	var repeating []byte
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			repeating = append(repeating, num[i])
		}
	}
	newArray := findUnique(repeating)
	return len(newArray) > 0
}

func findUnique(array []byte) []byte {
	newArray := []byte{}
	for i := range array {
		foundDuplicate := false
		for j := range array {
			if i != j && array[i] == array[j] {
				foundDuplicate = true
				break
			}
		}
		if !foundDuplicate {
			newArray = append(newArray, array[i])
		}
	}
	return newArray
}

func isSixDigit(num string) bool {
	return len(num) == 6
}

func isWithinRange(num, min, max int) bool {
	return num >= min && num <= max
}

func neverDecrease(num string) bool {
	decrease := false
	for i := 0; i < len(num)-1; i++ {
		firstNum, _ := strconv.Atoi(string(num[i]))
		secondNum, _ := strconv.Atoi(string(num[i+1]))
		if firstNum > secondNum {
			decrease = true
			break
		}
	}
	return !decrease
}

func isValid(num, min, max int) bool {
	numStr := strconv.Itoa(num)
	return isSixDigit(numStr) && isWithinRange(num, min, max) && neverDecrease(numStr) && twoAdjacentDigit(numStr) && twoAdjacentDigitNotPartOfLargerGroup(numStr)
}

func init() {
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	// inputArray := strings.Fields(string(inputByte))
	inputArray := strings.FieldsFunc(string(inputByte), func(r rune) bool {
		return r == '-'
	})
	// var err error
	min, err = strconv.Atoi(inputArray[0])
	if err != nil {
		panic(err)
	}

	max, err = strconv.Atoi(inputArray[1])
	if err != nil {
		panic(err)
	}
	// inputs, err = parseStringArrayToFloat64(inputArray)
	// if err != nil {
	// 	panic(err)
	// }
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
