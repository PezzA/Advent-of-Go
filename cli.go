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
	fmt.Println("	USEAGE: advent-of-go <year> <day>")
	fmt.Println("		<year> = year number of puzzle to run.  e.g. 2017")
	fmt.Println("		<day> = day number of puzzle to run.  e.g. 2")
}

func checkParams() (int, int, error) {
	if len(os.Args) < 2 {
		return 0, 0, errors.New("Not all parameters specified")
	}

	year, err := strconv.Atoi(os.Args[1])

	if err != nil {
		return 0, 0, errors.New("Called with invalid year")
	}

	day, err := strconv.Atoi(os.Args[2])

	if err != nil {
		return 0, 0, errors.New("Called with invalid day")
	}

	return year, day, nil
}
