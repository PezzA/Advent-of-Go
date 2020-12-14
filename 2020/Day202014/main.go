package Day202014

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type instructionType int

const (
	setMask instructionType = iota
	setAddress
)

type instruction struct {
	mask    string
	address int36
	value   int36
}

type int36 int64

func (i instruction) getType() instructionType {
	if i.mask != "" {
		return setMask
	}
	return setAddress
}

// get36String will return int36 as binary padded to 36 characters
func (i int36) get36String() string {
	raw := strconv.FormatInt(int64(i), 2)
	return fmt.Sprintf("%036s", raw)
}

func (i int36) applyMask(mask string) int36 {
	binary := i.get36String()

	newBinary := make([]byte, 0)
	for i := 0; i < len(mask); i++ {
		if mask[i] == 48 || mask[i] == 49 {
			newBinary = append(newBinary, mask[i])
			continue
		}

		newBinary = append(newBinary, binary[i])
	}

	return getInt(string(newBinary))
}

func (i int36) getAddresses(mask string) []int36 {
	binary := i.get36String()
	addresses := make([]int36, 0)
	newBinary := make([]byte, 0)

	xCount := make([]int, 0)

	// updated mask as detailed in step 2.  Ignore 0, apply X
	for i := 0; i < len(mask); i++ {
		if mask[i] == 88 || mask[i] == 49 {
			if mask[i] == 88 {
				xCount = append(xCount, i)
			}
			newBinary = append(newBinary, mask[i])
			continue
		}

		newBinary = append(newBinary, binary[i])
	}

	// will be 2 raised to the count of x permuations
	loop := int(math.Pow(2, float64(len(xCount))))

	// walk a padded binary sequence and replace on the x's
	for i := 0; i < loop; i++ {
		replaceMask := fmt.Sprintf("%0*b\n", len(xCount), i)

		permVal := newBinary
		for r := 0; r < len(xCount); r++ {
			permVal[xCount[r]] = replaceMask[r]
		}

		addresses = append(addresses, getInt(string(permVal)))
	}

	return addresses
}

func getInt(input string) int36 {
	output, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int36(output)
}

func getData(input string) []instruction {
	lines := strings.Split(input, "\n")
	data := make([]instruction, len(lines))

	for i := range lines {
		if lines[i][:4] == "mask" {
			var mask string
			fmt.Sscanf(lines[i], "mask = %s", &mask)
			data[i] = instruction{mask: mask}
		} else {
			var address, value int
			fmt.Sscanf(lines[i], "mem[%d] = %d", &address, &value)
			data[i] = instruction{mask: "", address: int36(address), value: int36(value)}
		}
	}
	return data

}
func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	mask := ""
	memory := make(map[int36]int36, 0)

	program := getData(Entry.PuzzleInput())

	for _, ins := range program {
		if ins.getType() == setMask {
			mask = ins.mask
			continue
		}

		memory[ins.address] = ins.value.applyMask(mask)
	}

	tot := int36(0)
	for _, v := range memory {
		tot += v
	}
	return fmt.Sprintf("%v", tot)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	mask := ""
	memory := make(map[int36]int36, 0)

	program := getData(Entry.PuzzleInput())

	for _, ins := range program {
		if ins.getType() == setMask {
			mask = ins.mask
			continue
		}

		addresses := ins.address.getAddresses(mask)
		for a := range addresses {
			memory[addresses[a]] = ins.value
		}

	}

	tot := int36(0)
	for _, v := range memory {
		tot += v
	}

	return fmt.Sprintf("%v", tot)
}
