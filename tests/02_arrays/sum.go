package main

// Sum adds a collection of numbers
func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}

	return
}

// SumAll adds all numbers within a collection from all collections
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}
