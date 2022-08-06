package cli

import (
	"fmt"
	"strconv"

	tm "github.com/buger/goterm"
)

var termWidthPrint = "%-" + strconv.Itoa(tm.Width()) + "v"

var partOnePreFix = tm.Color("    Part One --> ", tm.WHITE)
var partTwoPreFix = tm.Color("    Part Two --> ", tm.WHITE)

// GetHeader takes some input vars and returns a term formatted string that can be held by the called and
// passed into rendering the frame
func FormatHeader(year int, day int, title string) string {
	epilette := tm.Color("---", tm.YELLOW)
	dispYear := tm.Bold(tm.Color(fmt.Sprintf("%v", year), tm.YELLOW))
	dispTitle := tm.Bold(tm.Color(title, tm.WHITE))
	dispDay := tm.Bold(tm.Color(fmt.Sprintf("Day %v", day), tm.WHITE))
	return fmt.Sprintf("%v %v - %v : %v %v", epilette, dispYear, dispDay, dispTitle, epilette)
}

// HideCursor hides the console cursor
func HideCursor() {
	tm.Print("\033[?25l")
	tm.Flush()
}

// ShowCursor shows the console cursor
func ShowCursor() {
	tm.Print("\033[?25h")
	tm.Flush()
}

func FormatUpdate(input string) string {
	return tm.Color(input, tm.CYAN)
}

func FormatAnswer(input string) string {
	return tm.Bold(tm.Color(input, tm.GREEN))
}

// Interrupted handles if the console app is told to shut down, Ctrl+C et...
func Interrupted() {
	tm.Println(tm.Bold(tm.Color("    Why no finishings?", tm.RED)))
	ShowCursor()
}

func Render(text string) int {
	tm.Print(text)
	retVal := tm.CurrentHeight()
	tm.Flush()
	return retVal
}

// NewFrame get ready to draw a new frame, which basically entails moving the cursor back to the top
func NewFrame(rowsToMoveUp int) {
	tm.MoveCursorUp(rowsToMoveUp)
}
