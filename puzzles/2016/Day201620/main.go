package Day201620

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type iprange struct {
	min int
	max int
}

func loadData(input string) []iprange {
	lines := strings.Split(input, "\n")
	ipdata := make([]iprange, len(lines))

	for index, line := range lines {
		bits := strings.Split(line, "-")

		low, _ := strconv.Atoi(bits[0])
		high, _ := strconv.Atoi(bits[1])

		ipdata[index] = iprange{min: low, max: high}
	}

	return ipdata
}

func updateBoundaries(ip iprange, rangeList map[iprange]bool) {

	// is this block wholly contained?
	for k := range rangeList {
		if ip.min >= k.min && ip.max <= k.max {
			return
		}
	}

	// Are there any wholly contained blocks? if os remove them
	for k := range rangeList {
		if k.min >= ip.min && k.max <= ip.max {
			delete(rangeList, k)
		}
	}

	// is it contigous on the left
	for k := range rangeList {
		if ip.min == k.max+1 {
			delete(rangeList, k)
			ip.min = k.min
		}
	}

	// is it contigous on the right
	for k := range rangeList {
		if ip.max+1 == k.min {
			delete(rangeList, k)
			ip.min = k.min
		}
	}

	// does it overlap on the left
	for k := range rangeList {
		if ip.min >= k.min && ip.min <= k.max {
			delete(rangeList, k)
			ip.min = k.min
		}
	}

	// does it overlap on the left
	for k := range rangeList {
		if ip.max >= k.min && ip.max <= k.max {
			delete(rangeList, k)
			ip.max = k.max
		}
	}

	rangeList[ip] = true
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {

	data := loadData(inputData)

	compiledRanges := make(map[iprange]bool, 0)

	sort.Slice(data, func(i, j int) bool {
		return data[i].min < data[j].min
	})

	for i := range data {
		updateBoundaries(data[i], compiledRanges)
	}

	final := make([]iprange, 0)
	for iprange := range compiledRanges {
		final = append(final, iprange)
	}

	sort.Slice(final, func(i, j int) bool {
		return final[i].min < final[j].min
	})

	return fmt.Sprint(final[0].max + 1)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := loadData(inputData)

	compiledRanges := make(map[iprange]bool, 0)

	sort.Slice(data, func(i, j int) bool {
		return data[i].min < data[j].min
	})

	for i := range data {
		updateBoundaries(data[i], compiledRanges)
	}

	final := make([]iprange, 0)
	for iprange := range compiledRanges {
		final = append(final, iprange)
	}

	sort.Slice(final, func(i, j int) bool {
		return final[i].min < final[j].min
	})

	count := 0
	for i := 1; i < len(final); i++ {
		count = count + ((final[i].min - final[i-1].max) - 1)
	}
	return fmt.Sprint(count)
}
