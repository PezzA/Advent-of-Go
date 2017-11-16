package main

import (
	"errors"

	"github.com/pezza/AoC2017/Day201501"
	"github.com/pezza/AoC2017/Day201502"
	"github.com/pezza/AoC2017/Day201611"
	"github.com/pezza/AoC2017/TestDay"
)

func getPuzzle(day int) (dayData, dayRunner, error) {
	data := [...]dayData{TestDay.Entry, Day201611.Entry, Day201501.Entry, Day201502.Entry}
	runners := [...]dayRunner{TestDay.Entry, Day201611.Entry, Day201501.Entry, Day201502.Entry}

	runnerFound := false
	var dayrunner dayRunner
	for _, runner := range runners {
		if runner.Day() == day {
			runnerFound = true
			dayrunner = runner
		}
	}

	dataFound := false
	var daydata dayData
	for _, data := range data {
		if data.Day() == day {
			dataFound = true
			daydata = data
		}
	}

	if !runnerFound || !dataFound {
		return nil, nil, errors.New("Day specified has not been fully implemented yet")
	}

	return daydata, dayrunner, nil
}
