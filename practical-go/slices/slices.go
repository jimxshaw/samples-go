package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int
	fmt.Println("len", len(s)) // len is "nil safe"

	// you can compare a slice to nil
	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4]
	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:1000]) // will panic, out of range

	s3 = append(s3, 100)
	fmt.Printf("s3 (appended) = %#v\n", s3)
	fmt.Printf("s2 (appended) = %#v\n", s2)
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))

	var s4 []int
	// s4 := make([]int, 0, 1000) // Use this if you already know the capacity you want
	for i := 1; i <= 1000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Println("s4", len(s4), cap(s4))

	fmt.Println("concat: ", concat([]string{"A", "B"}, []string{"C", "D", "E"}))

	vs := []float64{4, 67, 3, 7, 4, 28}
	fmt.Println(median(vs))
	vs = []float64{4, 12, 7, 4, 28}
	fmt.Println(median(vs))
	fmt.Println(vs)
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)
	i := len(nums) / 2

	// if len(nums)&1 != 0 {
	if len(nums)%2 != 0 {
		return nums[i], nil
	}

	return (nums[i-1] + nums[i]) / 2, nil
}

func concat(s1, s2 []string) []string {
	// Restrictions: no "for" loops
	result := make([]string, len(s1)+len(s2))
	copy(result, s1)
	copy(result[len(s1):], s2)
	// append(s1, s2...) // this way works as well
	return result
}

func appendInt(s []int, v int) []int {
	i := len(s)
	// Enough space in underlying space.
	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		// Need to re-allocate and copy.
		fmt.Printf("re-allocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	s[i] = v
	return s
}
