package main

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	// Test cases including edge cases.
	testCases := []struct {
		inputSize int
		expected  [][]int
	}{
		{0, [][]int{}},
		{1, [][]int{{1}}},
		{2, [][]int{{1, 2}, {4, 3}}},
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{4, [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
	}

	for _, testCase := range testCases {
		actualResult := matrix(testCase.inputSize)
		if !matricesAreEqual(actualResult, testCase.expected) {
			t.Errorf("matrix(%d) = %v; expected %v", testCase.inputSize, actualResult, testCase.expected)
		}
	}
}

// matricesAreEqual checks if two 2D integer slices are equivalent.
func matricesAreEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
