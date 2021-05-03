package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	tm "github.com/buger/goterm"
)

var termWidthPrint = "%-" + strconv.Itoa(tm.Width()) + "v"

var partOnePreFix = tm.Color("    Part One --> ", tm.WHITE)
var partTwoPreFix = tm.Color("    Part Two --> ", tm.WHITE)

// OutputUseage outputs useage pattern
func outputUseage(errorMsg error) {
	fmt.Println(errorMsg)
	fmt.Println("")
	fmt.Println("	USEAGE: advent-of-code.exe [-vis] <year> <day>")
	fmt.Println("		<year> = year number of puzzle to run.  e.g. 2017")
	fmt.Println("		<day> = day number of puzzle to run.  e.g. 2")
	fmt.Println("		")
}

// CheckParams checks input parameters are correct
func checkParams() (int, int, error) {

	if len(os.Args) < 3 {
		return 0, 0, errors.New("not all parameters specified")
	}

	year, err := strconv.Atoi(os.Args[1])

	if err != nil {
		return 0, 0, errors.New("called with invalid year")
	}

	day, err := strconv.Atoi(os.Args[2])

	if err != nil {
		return 0, 0, errors.New("called with invalid day")
	}

	return year, day, nil
}

// GetHeader takes some input vars and returns a term formatted string that can be held by the called and
// passed into rendering the frame
func formatHeader(year int, day int, title string) string {
	epilette := tm.Color("---", tm.YELLOW)
	dispYear := tm.Bold(tm.Color(fmt.Sprintf("%v", year), tm.YELLOW))
	dispTitle := tm.Bold(tm.Color(title, tm.WHITE))
	dispDay := tm.Bold(tm.Color(fmt.Sprintf("Day %v", day), tm.WHITE))
	return fmt.Sprintf("%v %v - %v : %v %v", epilette, dispYear, dispDay, dispTitle, epilette)
}

// HideCursor hides the console cursor
func hideCursor() {
	tm.Print("\033[?25l")
	tm.Flush()
}

// ShowCursor shows the console cursor
func showCursor() {
	tm.Print("\033[?25h")
	tm.Flush()
}

func formatUpdate(input string) string {
	return tm.Color(input, tm.CYAN)
}

func formatAnswer(input string) string {
	return tm.Bold(tm.Color(input, tm.GREEN))
}

// Interrupted handles if the console app is told to shut down, Ctrl+C et...
func interrupted() {
	tm.Println(tm.Bold(tm.Color("    Why no finishings?", tm.RED)))
	showCursor()
}

func render(text string) int {
	tm.Print(text)
	retVal := tm.CurrentHeight()
	tm.Flush()
	return retVal
}

// NewFrame get ready to draw a new frame, which basically entails moving the cursor back to the top
func newFrame(rowsToMoveUp int) {
	tm.MoveCursorUp(rowsToMoveUp)
}
