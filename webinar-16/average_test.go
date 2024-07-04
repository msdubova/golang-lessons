package main

import (
	"bytes"
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

func TestWriteAverage(t *testing.T) {
	elements := []int{1, 2, 3}
	average := 2
	expectedOutput := `{"average":2,"elements":[1,2,3]}`

	var buff = &bytes.Buffer{}

	err := WriteAverage(buff, elements, average)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if out := buff.String(); out != expectedOutput {
		t.Fatalf("Got unexpected output: %v", out)
	}
}
