package main

import "math/rand"

func GenerateLotsOfRandomNumbers() []int {
	nums := make([]int, 10_000)

	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	return nums
}

func GenerateLotsOfRandomNumbersConcurrent() []int {
	ch := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				ch <- rand.Intn(1000)
			}
		}()
	}

	nums := make([]int, 10_000)
	var i int

	for {
		num := <-ch
		nums[i] = num
		i++
		if i >= 10_000 {
			break
		}
	}

	return nums
}
