package main

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput int
	}{
		{
			input:          []int{1, 2, 3, 10},
			expectedOutput: 4,
		},
		{
			input:          []int{10},
			expectedOutput: 10,
		},
		{
			input:          []int{100, 200, 3, 10},
			expectedOutput: 78,
		},
		{
			input:          []int{-100, 200, -3, 10},
			expectedOutput: 26,
		},
		{
			input:          []int{},
			expectedOutput: 0,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actualOutput := Average(test.input)
			if actualOutput != test.expectedOutput {
				t.Fatalf("Got unexpected result: %v. Want: %v", actualOutput, test.expectedOutput)
			}
		})
	}
}
