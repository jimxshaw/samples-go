package main

// Sum adds a collection of numbers
func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}

	return
}
