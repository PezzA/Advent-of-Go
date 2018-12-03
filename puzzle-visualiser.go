package main

import (
	"fmt"
	"time"
)

func visualise(puzzle visualiser) {
	updateChan := make(chan []string)
	go puzzle.Visualise(updateChan)

	for update := range updateChan {
		for _, line := range update {
			fmt.Println(line)
		}
		time.Sleep(time.Millisecond * 17)
	}
}
