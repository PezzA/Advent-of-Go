package main

import (
	"errors"

	"github.com/pezza/AoC2017/Day201501"
	"github.com/pezza/AoC2017/Day201502"
	"github.com/pezza/AoC2017/Day201503"
	"github.com/pezza/AoC2017/Day201504"
	"github.com/pezza/AoC2017/Day201505"
	"github.com/pezza/AoC2017/Day201611"
	"github.com/pezza/AoC2017/Day201617"
	"github.com/pezza/AoC2017/Day201619"
	"github.com/pezza/AoC2017/Day201622"
	"github.com/pezza/AoC2017/Day201701"
	"github.com/pezza/AoC2017/Day201702"
	"github.com/pezza/AoC2017/Day201703"
	"github.com/pezza/AoC2017/TestDay"
)

func getPuzzle(day int) (dayData, dayRunner, error) {
	data := [...]dayData{
		TestDay.Entry,

		Day201501.Entry,
		Day201502.Entry,
		Day201503.Entry,
		Day201504.Entry,
		Day201505.Entry,
		Day201611.Entry,
		Day201617.Entry,
		Day201619.Entry,
		Day201622.Entry,
		Day201701.Entry,
		Day201702.Entry,
		Day201703.Entry,
	}

	runners := [...]dayRunner{
		TestDay.Entry,
		Day201501.Entry,
		Day201502.Entry,
		Day201503.Entry,
		Day201504.Entry,
		Day201505.Entry,
		Day201611.Entry,
		Day201617.Entry,
		Day201619.Entry,
		Day201622.Entry,
		Day201701.Entry,
		Day201702.Entry,
		Day201703.Entry,
	}

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
