package Day201504

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2015, 04, "The Ideal Stocking Stuffer", 2
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	hasher := md5.New()

	i := 0
	for {
		hasher.Reset()
		hasher.Write([]byte(inputData + strconv.Itoa(i)))

		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash[:5] == "00000" {
			return strconv.Itoa(i)
		}
		i++
	}
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	hasher := md5.New()

	i := 0
	for {
		hasher.Reset()
		hasher.Write([]byte(inputData + strconv.Itoa(i)))

		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash[:6] == "000000" {
			return strconv.Itoa(i)
		}
		i++
	}
}

func (td dayEntry) PuzzleInput() string {
	return "iwrupvqb"
}
