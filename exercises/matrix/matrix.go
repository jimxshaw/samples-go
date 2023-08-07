package main

import "fmt"

// matrix generates an NxN spiral matrix.
func matrix(n int) [][]int {
	// Initialize the matrix with zeros.
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Define boundaries for the spiral progression.
	upperRow, lowerRow, leftCol, rightCol := 0, n-1, 0, n-1
	currentValue := 1

	// Continue filling the matrix while the currentValue is less than or equal to n^2.
	for currentValue <= n*n {
		// Fill from left to right in the uppermost unfilled row.
		for col := leftCol; col <= rightCol && currentValue <= n*n; col++ {
			matrix[upperRow][col] = currentValue
			currentValue++
		}
		upperRow++

		// Fill from top to bottom in the rightmost unfilled column.
		for row := upperRow; row <= lowerRow && currentValue <= n*n; row++ {
			matrix[row][rightCol] = currentValue
			currentValue++
		}
		rightCol--

		// Fill from right to left in the lowermost unfilled row.
		for col := rightCol; col >= leftCol && currentValue <= n*n; col-- {
			matrix[lowerRow][col] = currentValue
			currentValue++
		}
		lowerRow--

		// Fill from bottom to top in the leftmost unfilled column.
		for row := lowerRow; row >= upperRow && currentValue <= n*n; row-- {
			matrix[row][leftCol] = currentValue
			currentValue++
		}
		leftCol++
	}

	return matrix
}

func main() {
	resultMatrix := matrix(4)
	for _, row := range resultMatrix {
		fmt.Println(row)
	}
}
