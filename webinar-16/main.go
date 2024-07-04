package main

import "os"

func main() {
	elements := []int{1, 2, 3, 4, 5}

	average := Average(elements)

	f, err := os.Create("average.json")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	err = WriteAverage(f, elements, average)
	if err != nil {
		panic(err.Error())
	}
}
