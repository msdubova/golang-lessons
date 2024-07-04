package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func WriteAverageToFile(elems []int, average int) error {
	content := map[string]any{
		"elements": elems,
		"average":  average,
	}

	rawContent, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("marshaling content: %w", err)
	}

	err = os.WriteFile("average.json", rawContent, os.ModePerm)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}
