package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
)

var termWidthPrint = "%-" + strconv.Itoa(tm.Width()) + "v"

var partOnePreFix = tm.Color("    Part One --> ", tm.WHITE)
var partTwoPreFix = tm.Color("    Part Two --> ", tm.WHITE)

// OutputUseage outputs useage pattern
func OutputUseage(errorMsg error) {
	fmt.Println(errorMsg)
	fmt.Println("")
	fmt.Println("	USEAGE: AoC2017 <year> <day> [<part>]")
	fmt.Println("		<year> = year number of puzzle to run.  e.g. 2017")
	fmt.Println("		<day> = day number of puzzle to run.  e.g. 2")
}

// CheckParams checks input parameters are correct
func CheckParams() (int, int, error) {
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

// GetHeader takes some input vars and returns a term formatted string that can be held by the called and
// pased into rendering the frame
func GetHeader(year int, day int, title string) string {
	epilette := tm.Color("---", tm.YELLOW)
	dispYear := tm.Bold(tm.Color(fmt.Sprintf("%v", year), tm.YELLOW))
	dispTitle := tm.Bold(tm.Color(title, tm.WHITE))
	dispDay := tm.Bold(tm.Color(fmt.Sprintf("Day %v", day), tm.WHITE))
	return fmt.Sprintf("%v %v - %v : %v %v", epilette, dispYear, dispDay, dispTitle, epilette)
}

// HideCursor hides the console cursor
func HideCursor() {
	tm.Print("\033[?25l")
}

// ShowCursor hides the console cursor
func ShowCursor() {
	tm.Print("\033[?25h")
	tm.Flush()
}

// Interrupted handles if the console app is told to shut down, Ctrl+C et...
func Interrupted() {
	tm.Println(tm.Bold(tm.Color("    Why no finishings?", tm.RED)))
	ShowCursor()
}

// DrawFrame takes a bunch of parameters that form the basic view model, then outputs to console
func DrawFrame(partOneResult string, partOneUpdate []string, partOneDuration time.Duration, partTwoResult string, partTwoUpdate []string, partTwoDuration time.Duration, header string) {
	tm.Println()
	tm.Println(header)
	tm.Println()
	if partOneResult == "" {
		if len(partOneUpdate) == 1 {
			drawUpdate(partOnePreFix, partOneUpdate[0])
		}
	} else {
		drawResult(partOnePreFix, partOneResult, partOneDuration)
	}

	if partTwoResult == "" {
		if len(partTwoUpdate) == 1 {
			drawUpdate(partTwoPreFix, partTwoUpdate[0])
		}
	} else {
		drawResult(partTwoPreFix, partTwoResult, partTwoDuration)
	}
	tm.Flush()

}

// NewFrame get ready to draw a new frame, which basically entails moving the cursor back to the top
func NewFrame(rowsToMoveUp int) {
	tm.MoveCursorUp(rowsToMoveUp)
}

func drawResult(prefix string, result string, timeTaken time.Duration) {
	update := tm.Bold(tm.Color(result, tm.GREEN))
	time := tm.Bold(tm.Color(fmt.Sprintf("The call took %v to run.", timeTaken), tm.WHITE))
	tm.Println(prefix, fmt.Sprintf(termWidthPrint, update+"    "+time))
}

func drawUpdate(prefix string, update string) {
	updateText := tm.Bold(tm.Color(update, tm.CYAN))
	tm.Println(prefix, fmt.Sprintf(termWidthPrint, updateText))
}
