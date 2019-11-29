package Day201715

import (
	"fmt"
	"strconv"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2017, 15, "Dueling Generators"
}

func lowest16bits(input int64) string {
	binString := strconv.FormatInt(input, 2)
	binString = fmt.Sprintf("%016s", binString)
	return string(binString[len(binString)-16:])
}

func nextFactor(factor int64, seed int64, requiredDivisor int64) int64 {
	val := (seed * factor) % 2147483647

	if requiredDivisor > 0 {

		valid := val%requiredDivisor == 0

		for !valid {
			val = (val * factor) % 2147483647
			valid = val%requiredDivisor == 0
		}
	}
	return val
}

func nextFactorChan(factor int64, seed int64, requiredDivisor int64, report chan int64) {
	i := 0
	for {
		seed = nextFactor(factor, seed, requiredDivisor)
		i++
		report <- seed
	}
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	genASeed, genAFactor := int64(873), int64(16807)
	genBSeed, genBFactor := int64(583), int64(48271)

	total := 0

	for i := 0; i < 40000000; i++ {
		genASeed = nextFactor(genAFactor, genASeed, 0)
		genBSeed = nextFactor(genBFactor, genBSeed, 0)

		if lowest16bits(genASeed) == lowest16bits(genBSeed) {
			total++
		}

		if i%10000 == 0 && updateChan != nil {
			updateChan <- []string{fmt.Sprintf("Compared %v of %v matches", i, 40000000)}
		}
	}
	return strconv.Itoa(total)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	genASeed, genAFactor := int64(873), int64(16807)
	genBSeed, genBFactor := int64(583), int64(48271)

	total := 0

	for i := 0; i < 5000000; i++ {
		genASeed = nextFactor(genAFactor, genASeed, 4)
		genBSeed = nextFactor(genBFactor, genBSeed, 8)

		if lowest16bits(genASeed) == lowest16bits(genBSeed) {
			total++
		}
		if i%1 == 0 && updateChan != nil {
			updateChan <- []string{fmt.Sprintf("Compared %v of %v matches", i, 40000000)}
		}
	}

	return strconv.Itoa(total)
}

func (td dayEntry) PartThree(inputData string) string {
	genASeed, genAFactor := int64(873), int64(16807)
	genBSeed, genBFactor := int64(583), int64(48271)

	genAChan := make(chan int64, 100)
	genBChan := make(chan int64, 100)

	total := 0

	go nextFactorChan(genAFactor, genASeed, 4, genAChan)
	go nextFactorChan(genBFactor, genBSeed, 8, genBChan)

	for i := 0; i < 5000000; i++ {

		genASeed = <-genAChan
		genBSeed = <-genBChan

		if lowest16bits(genASeed) == lowest16bits(genBSeed) {
			total++
		}
	}
	return strconv.Itoa(total)
}

func (td dayEntry) PuzzleInput() string {
	return "873,583"
}
