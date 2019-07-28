package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstTwice(t *testing.T) {
	t.Run("+1, -1", func(t *testing.T) {
		actual := firstTwice([]int{1, -1})
		expected := 0
		assert.Equal(t, expected, actual, "expected %v actual %v", expected, actual)
	})
	t.Run("+3, +3, +4, -2, -4", func(t *testing.T) {
		actual := firstTwice([]int{3, 3, 4, -2, -4})
		expected := 10
		assert.Equal(t, expected, actual, "expected %v actual %v", expected, actual)
	})
	t.Run("-6, +3, +8, +5, -6", func(t *testing.T) {
		actual := firstTwice([]int{-6, 3, 8, 5, -6})
		expected := 5
		assert.Equal(t, expected, actual, "expected %v actual %v", expected, actual)
	})
	t.Run("+7, +7, -2, -7, -4", func(t *testing.T) {
		actual := firstTwice([]int{7, 7, -2, -7, -4})
		expected := 14
		assert.Equal(t, expected, actual, "expected %v actual %v", expected, actual)
	})
}
