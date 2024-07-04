package main

import "testing"

func TestAverage(t *testing.T) {
	input := []int{1, 2, 3, 10}
	expectedOutput := 4

	actualOutput := Average(input)
	if actualOutput != expectedOutput {
		t.Fatalf("Got unexpected result: %v. Want: %v", actualOutput, expectedOutput)
	}
}
