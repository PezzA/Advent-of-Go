package puzzles

type PuzzlePart func(string, chan []string) string

type Contenter interface {
	Describe() (int, int, string, int)
	PuzzleInput() string
}

type Reporter interface {
	State() int
}

type DailyPuzzle interface {
	Contenter
	PartOne(string, chan []string) string
	PartTwo(string, chan []string) string
}

type PuzzleState int

const (
	Unsolved PuzzleState = iota
	First
	Second
	Other
)
