package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func findAnswer() int {
	var total int
	for _, item := range unionArray {
		// fmt.Println(item)
		total = total + findDistanceToCOM(left, right, item)
	}
	return total
}

func findDistanceToCOM(left, right []string, item string) int {
	if item == "COM" {
		return 0
	}
	index := findIndex(right, item)
	if index < 0 {
		panic(fmt.Sprintf("index not found for %s", item))
	}
	parent := findParent(left, index)
	return 1 + findDistanceToCOM(left, right, parent)
}

func findParent(array []string, index int) string {
	return array[index]
}

func findIndex(array []string, item string) int {
	for i, value := range array {
		// fmt.Printf("Comparing %s with %s\n", value, item)
		if value == item {
			return i
		}
	}
	return -1
}
func createTree() {

}
func main() {
	start := time.Now()
	answer := findAnswer()
	fmt.Printf("The answer is: %+v. That took %s\n", int(answer), time.Since(start))
}

var input string
var unionArray []string
var left []string
var right []string

func init() {
	var err error
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputString := string(inputByte)
	inputArray := strings.Fields(inputString)

	for _, item := range inputArray {
		orbitMap := strings.FieldsFunc(item, func(r rune) bool {
			return r == ')'
		})
		left = append(left, orbitMap[0])
		right = append(right, orbitMap[1])
	}
	unionArray = Union(left, right)

	// fmt.Println(inputArray)
	// fmt.Println(left)
	// fmt.Println(right)
	// fmt.Println(unionArray)
}

func Union(left, right []string) []string {
	var newArray []string
	uniqueMap := make(map[string]bool)

	for i := 0; i < len(left); i++ {
		if _, ok := uniqueMap[left[i]]; !ok {
			uniqueMap[left[i]] = true
			newArray = append(newArray, left[i])
		}
	}

	for j := 0; j < len(right); j++ {
		if _, ok := uniqueMap[right[j]]; !ok {
			uniqueMap[right[j]] = true
			newArray = append(newArray, right[j])
		}
	}
	return newArray
}
