package main

func main() {
	elements := []int{1, 2, 3, 4, 5}

	average := Average(elements)

	err := WriteAverageToFile(elements, average)
	if err != nil {
		panic(err.Error())
	}
}
