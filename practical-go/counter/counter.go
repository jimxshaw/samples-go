package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0

	const n = 10

	var wg sync.WaitGroup

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			wg.Done()
			for j := 0; j < 10_000; j++ {
				count++
			}
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
