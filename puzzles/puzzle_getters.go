package puzzles

import (
	"errors"
	"math"
)

func GetPuzzle(day int, year int) (DailyPuzzle, error) {

	for _, puzzle := range puzzleList {
		puzzleYear, puzzleDay, _, _ := puzzle.Describe()
		if puzzleDay == day && puzzleYear == year {
			return puzzle, nil
		}
	}

	return nil, errors.New("day specified has not been fully implemented yet")
}

func GetYearRange() (int, int) {
	min, max := math.MaxInt32, 0

	for _, puzzle := range puzzleList {
		puzzleYear, _, _, _ := puzzle.Describe()

		if puzzleYear > max {
			max = puzzleYear
		}

		if puzzleYear < min && puzzleYear != 0 {
			min = puzzleYear
		}
	}
	return min, max
}

func GetManualYearRange() (int, int) {
	return 2015, 2022
}
