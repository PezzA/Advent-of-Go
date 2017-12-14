package Day201710

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	common "github.com/pezza/AoC2017/Common"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry testDay

type testDay bool

func getElements(size int) []int {
	elements := make([]int, 0)

	for i := 0; len(elements) < size; i++ {
		elements = append(elements, i)
	}

	return elements
}

var standardSize = 256
var lengthSuffix = []byte{17, 31, 73, 47, 23}

func swap(elements []int, length int, currentPos int, skipSize int) ([]int, int, int) {
	swaps := make([]int, 0)
	aft := make([]int, 0)

	for i := 0; i <= len(elements); i++ {
		index := common.GetWrappedIndex(currentPos+i, len(elements), 0)
		if len(swaps) == length {
			aft = append(aft, elements[index])
		} else {
			swaps = append(swaps, elements[index])
		}
	}

	swappedSwaps := make([]int, 0)
	for i := len(swaps) - 1; i >= 0; i-- {
		swappedSwaps = append(swappedSwaps, swaps[i])
	}

	newOrder := append(swappedSwaps, aft...)

	for i := 0; i < len(elements); i++ {
		index := common.GetWrappedIndex(currentPos+i, len(elements), 0)
		elements[index] = newOrder[i]
	}

	newIndex := common.GetWrappedIndex(currentPos, len(elements), length+skipSize)

	return elements, newIndex, skipSize + 1
}

func doSwaps(lengths []int, elements []int, currentPos int, skipSize int) ([]int, int, int) {
	for _, val := range lengths {
		elements, currentPos, skipSize = swap(elements, val, currentPos, skipSize)
	}

	return elements, currentPos, skipSize
}

func DoFullHash(inputs string) string {
	byteArray := append([]byte(inputs), lengthSuffix...)

	lengths := make([]int, 0)

	for _, val := range byteArray {
		lengths = append(lengths, int(val))
	}

	currentPos := 0
	swapSize := 0

	var elements []int
	constElements := getElements(standardSize)

	for i := 0; i < 64; i++ {
		elements, currentPos, swapSize = doSwaps(lengths, constElements, currentPos, swapSize)
	}

	var buffer bytes.Buffer
	for i := 0; i < len(elements); i += 16 {
		val := getDenseHashSection(elements[i : i+16])
		buffer.WriteString(fmt.Sprintf("%02x", val))
	}

	return buffer.String()
}

func getDenseHashSection(bits []int) int {
	return bits[0] ^ bits[1] ^ bits[2] ^ bits[3] ^ bits[4] ^ bits[5] ^ bits[6] ^ bits[7] ^ bits[8] ^ bits[9] ^ bits[10] ^ bits[11] ^ bits[12] ^ bits[13] ^ bits[14] ^ bits[15]
}

func (td testDay) PartOne(inputData string) (string, error) {
	bits := strings.Split(inputData, ",")

	lengths := make([]int, 0)

	for _, val := range bits {
		intVal, _ := strconv.Atoi(val)

		lengths = append(lengths, intVal)
	}

	elements, _, _ := doSwaps(lengths, getElements(standardSize), 0, 0)

	return strconv.Itoa(elements[0] * elements[1]), nil
}

func (td testDay) PartTwo(inputData string) (string, error) {

	return DoFullHash(inputData), nil
}

func (td testDay) Day() int {
	return 10
}

func (td testDay) Year() int {
	return 2017
}

func (td testDay) Title() string {
	return "Knot Hash"
}

func (td testDay) GetData() string {
	return "14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244"
}
