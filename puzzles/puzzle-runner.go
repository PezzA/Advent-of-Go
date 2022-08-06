package puzzles

import (
	"fmt"
	"time"

	"github.com/pezza/advent-of-code/cli"
)

type partResult struct {
	answered bool
	result   string
	time     time.Duration
}

func Runner(puzzle DailyPuzzle) {
	cli.HideCursor()

	defer cli.ShowCursor()

	inputData := puzzle.PuzzleInput()
	year, day, title := puzzle.Describe()

	header := cli.FormatHeader(year, day, title)

	part1Answer, part2Answer := partResult{false, "", 0}, partResult{false, "", 0}
	part1Update, part2Update := []string{"*"}, []string{"*"}

	part1Chan := make(chan partResult)
	part2Chan := make(chan partResult)

	part1UpdateChan := make(chan []string, 10)
	part2UpdateChan := make(chan []string, 10)

	currline := 0

	go doPart(puzzle.PartOne, inputData, part1Chan, part1UpdateChan)
	go doPart(puzzle.PartTwo, inputData, part2Chan, part2UpdateChan)

	complete := false

	for !complete {
		select {
		case result := <-part1Chan:
			part1Answer = result
		case result := <-part2Chan:
			part2Answer = result
		case update := <-part1UpdateChan:
			part1Update = update
		case update := <-part2UpdateChan:
			part2Update = update
		default:
			cli.NewFrame(currline)

			p1Update, p2Update := cli.FormatUpdate(part1Update[0]), cli.FormatUpdate(part2Update[0])

			if part1Answer.answered {
				p1Update = formatPart(part1Answer)
			}

			if part2Answer.answered {
				p2Update = formatPart(part2Answer)
			}

			currline = cli.Render(getTextToRender(header, p1Update, p2Update))

			if part1Answer.answered && part2Answer.answered {
				complete = true
			}

			time.Sleep(17 * time.Millisecond)
		}
	}

	if year == 2018 && day == 10 {
		fmt.Println(part1Answer.result)
	}

	if year == 2019 && day == 8 {
		fmt.Println(part2Answer.result)
	}
}

func getTextToRender(header, partOneUpdate, partTwoUpdate string) string {
	return fmt.Sprintf("\n%v\n    Part One --> %v\n    Part Two --> %v\n", header, partOneUpdate, partTwoUpdate)
}

func formatPart(res partResult) string {
	return fmt.Sprintf("%v (%v)", cli.FormatAnswer(res.result), res.time)
}

func doPart(fn PuzzlePart, inputData string, response chan partResult, updateChan chan []string) {
	start := time.Now()
	output := fn(inputData, updateChan)
	end := time.Now()
	elapsed := end.Sub(start)
	response <- partResult{true, output, elapsed}
}
