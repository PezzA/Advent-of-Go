package Day201504

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

func (td testDay) PartOne(inputData string) (string, error) {
	hasher := md5.New()

	i := 0
	for {
		hasher.Reset()
		hasher.Write([]byte(inputData + strconv.Itoa(i)))

		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash[:5] == "00000" {
			return strconv.Itoa(i), nil
		}
		i++
	}
}

func (td testDay) PartTwo(inputData string) (string, error) {
	hasher := md5.New()

	i := 0
	for {
		hasher.Reset()
		hasher.Write([]byte(inputData + strconv.Itoa(i)))

		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash[:6] == "000000" {
			return strconv.Itoa(i), nil
		}
		i++
	}
}

func (td testDay) Day() int {
	return 201504
}

func (td testDay) Heading() string {
	return "--- (2015) Day 4: The Ideal Stocking Stuffer ---"
}

func (td testDay) GetTestData() ([]string, []string) {
	return []string{
			"abcdef",
			"pqrstuv",
		},
		[]string{}
}

func (td testDay) GetData() string {
	return "iwrupvqb"
}
