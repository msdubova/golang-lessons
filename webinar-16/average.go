package main

func Average(elems []int) int {
	var sum int
	for _, e := range elems {
		sum += e
	}
	return sum / len(elems)
}
