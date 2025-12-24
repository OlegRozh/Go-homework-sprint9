package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestGenerateRandomElementsLength(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 50, expected: 50},
		{input: 1000, expected: 1000},
	}
	for _, test := range tests {
		result := generateRandomElements(test.input)
		assert.Len(t, result, test.expected,
			"generateRandomElements(%d) should return slice with length %d",
			test.input, test.expected)
	}
}

func TestGenerateRandomElementsEmpty(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"negative size", -1},
		{"zero size", 0},
	}
	for _, test := range tests {
		result := generateRandomElements(test.size)
		assert.NotNil(t, result, "The result should not be nil")
		assert.Empty(t, result, "\"generateRandomElements(%d) should return empty slice\", tc.size")
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	result := maximum([]int{})
	assert.Equal(t, 0, result, "The maximum slice should be 0")
}

func TestMaximumSingleElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{50}, 50},
		{[]int{-5}, -5},
		{[]int{0}, 0},
	}
	for _, v := range tests {
		result := maximum(v.input)
		assert.Equal(t, v.expected, result, "The maximum slice should be equal")
	}
}

func TestMaximumSomeElements(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 5},
	}
	for _, v := range tests {
		result := maximum(v.input)
		assert.Equal(t, v.expected, result, "The maximum slice should be equal")
	}
}

func TestMaximumDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{7, 7, 7}, 7},
		{[]int{-3, -3, -3}, -3},
	}
	for _, v := range tests {
		result := maximum(v.input)
		assert.Equal(t, v.expected, result, "The maximum slice should be equal")
	}
}

func TestMaximumNegative(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{-1, -2, -3, -4, -5}, -1},
		{[]int{-10, -5, -8}, -5},
	}
	for _, v := range tests {
		result := maximum(v.input)
		assert.Equal(t, v.expected, result, "The maximum slice should be equal")
	}
}
