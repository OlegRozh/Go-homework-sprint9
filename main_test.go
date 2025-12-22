package main

import (
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElementsLength(t *testing.T) {
	testSize := 5
	result := generateRandomElements(testSize)
	if len(result) != testSize {
		t.Errorf("returned wrong length: got %d, want %d", len(result), testSize)
	}
}

func TestGenerateRandomElementsEmpty(t *testing.T) {
	result := generateRandomElements(-1)
	if len(result) != 0 {
		t.Errorf("If the size is negative, an empty slice should be returned")
	}
	result = generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("If the size is zero, an empty slice should be returned")
	}
}

func TestGenerateRandomElementsRange(t *testing.T) {
	testSize := 1000
	result := generateRandomElements(testSize)

	for i, v := range result {
		if v < 0 || v >= SIZE {
			t.Errorf("element at index %d: value %d is out of range [0, %d)",
				i, v, SIZE)
		}
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Errorf("Maximum slice is empty, an empty slice should be returned")
	}
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
		if result != v.expected {
			t.Errorf("Maximum slice is %d, want %d", result, v.expected)
		}
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
		if result != v.expected {
			t.Errorf("Maximum slice is %d, want %d", result, v.expected)
		}
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
		if result != v.expected {
			t.Errorf("Maximum slice is %d, want %d", result, v.expected)
		}
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
		if result != v.expected {
			t.Errorf("Maximum slice is %d, want %d", result, v.expected)
		}
	}
}
