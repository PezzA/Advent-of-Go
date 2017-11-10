package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func outputUseage(errorMsg error) {
	fmt.Println(errorMsg)
	fmt.Println("")
	fmt.Println("	USEAGE: advent-of-go <day>")
	fmt.Println("		<day> = day number of puzzle to run.  e.g. 2")
	fmt.Println("")
	fmt.Println("	Flags")
	fmt.Println("		--test : run the day with test data")
}

func checkParams() (int, bool, error) {
	if len(os.Args) < 2 {
		return 0, false, errors.New("Not all parameters specified")
	}

	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		return 0, false, errors.New("Called with invalid parameters")
	}

	runTest := false
	for _, val := range os.Args {
		if val == "--test" {
			runTest = true
		}
	}

	return day, runTest, nil
}
