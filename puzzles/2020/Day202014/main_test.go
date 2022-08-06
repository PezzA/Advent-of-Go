package Day202014

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	Expect(data[0].getType()).Should(Equal(setMask))
	Expect(data[0].mask).Should(Equal("100X000X100X00XX1010X0001X11XX100110"))
	Expect(data[1].getType()).Should(Equal(setAddress))
	Expect(data[1].address).Should(Equal(int36(33470)))
	Expect(data[1].value).Should(Equal(int36(43619)))
}

func Test_Get36String(t *testing.T) {
	RegisterTestingT(t)

	Expect(int36(11).get36String()).Should(Equal("000000000000000000000000000000001011"))
	Expect(getInt("1011")).Should(Equal(int36(11)))
	Expect(int36(101).get36String()).Should(Equal("000000000000000000000000000001100101"))
	Expect(getInt("1100101")).Should(Equal(int36(101)))
	Expect(int36(64).get36String()).Should(Equal("000000000000000000000000000001000000"))
	Expect(getInt("1000000")).Should(Equal(int36(64)))
}

func Test_ApplyMask(t *testing.T) {
	RegisterTestingT(t)

	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"

	Expect(int36(11).applyMask(mask)).Should(Equal(int36(73)))
	Expect(int36(101).applyMask(mask)).Should(Equal(int36(101)))
	Expect(int36(0).applyMask(mask)).Should(Equal(int36(64)))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

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

	Expect(tot).Should(Equal(int36(165)))

}

func Test_getAddresses(t *testing.T) {
	RegisterTestingT(t)

	mask := "000000000000000000000000000000X1001X"
	fmt.Println(int36(42).getAddresses(mask))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

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

	Expect(tot).Should(Equal(int36(208)))

}

func Benchmark_BenchPartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data, nil)
	}
}

func Benchmark_BenchPartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data, nil)
	}
}
