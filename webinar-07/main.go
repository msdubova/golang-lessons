package main

import (
	"fmt"
	"math/rand/v2"
	"webinar07/batch"
	"webinar07/log"
)

// Маємо певний тип LogInserter - він записує логи.
// Ціль - зробити LogInserter конкурентним та відправляти дані раз в N секунд.

func main() {
	var inserter log.Inserter

	batchInserter := batch.NewInserter(inserter)

	generateAndInsertLogs(batchInserter)

	// TODO: add close() to save all the data
}

type Inserter interface {
	Insert(logs []log.Log)
}

func generateAndInsertLogs(inserter Inserter) {
	for i := 0; i < 10; i++ {
		logs := make([]log.Log, 0, rand.IntN(10)+1)

		for j := 0; j < cap(logs); j++ {
			l := log.NewRandom(fmt.Sprintf("%d-%d", i, j))

			logs = append(logs, l)
		}

		inserter.Insert(logs)
	}
}

func exampleGoTo() {
LABEL_1:
	fmt.Println("Hello")

	goto LABEL_1
}
