package Day201617

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

type directions struct {
	startingSeed string
	up           bool
	down         bool
	left         bool
	right        bool
	hash         string
	x            int
	y            int
}

func traverse(input string, startx int, starty int) directions {

	if startx == 4 && starty == 4 {
		// we have found the vault
	}

	hasher := md5.New()

	hasher.Reset()
	hasher.Write([]byte(input))

	hash := hex.EncodeToString(hasher.Sum(nil))

	room := directions{
		startingSeed: input,
		up:           isOpen(hash[0]) && starty > 1,
		down:         isOpen(hash[1]) && starty < 4,
		left:         isOpen(hash[2]) && startx > 1,
		right:        isOpen(hash[3]) && startx < 4,
		x:            startx,
		y:            starty,
		hash:         hash,
	}

	if startx != 4 && starty != 4 {
		// we have found the vault
	}

	return room
}

func isOpen(input byte) bool {
	return input == 'b' || input == 'c' || input == 'd' || input == 'e' || input == 'f'
}

func (td testDay) PartOne(inputData string) (string, error) {

	fmt.Println(traverse(inputData, 1, 1))

	return "", nil

}

func (td testDay) PartTwo(inputData string) (string, error) {
	return " -- Not Yet Implemented --", nil
}

func (td testDay) Day() int {
	return 201617
}

func (td testDay) Heading() string {
	return "--- (2016) Day 17: Two Steps Forward ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{
			"hijkl",
			//"ihgpwlah",
			//"kglvqrro",
			//"ulqzkmiv",
		},
		[]string{""}
}

func (td testDay) GetData() string {
	return "ioramepc"
}
