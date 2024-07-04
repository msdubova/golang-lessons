package main

func Average(elems []int) int {
	if len(elems) == 0 {
		return 0
	}

	var sum int
	for _, e := range elems {
		sum += e
	}

	return sum / len(elems)
}
