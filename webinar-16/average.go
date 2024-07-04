package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func WriteAverage(w io.Writer, elems []int, average int) error {
	content := map[string]any{
		"elements": elems,
		"average":  average,
	}

	rawContent, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("marshaling content: %w", err)
	}

	_, err = w.Write(rawContent)
	if err != nil {
		return fmt.Errorf("writing average: %w", err)
	}

	return nil
}
