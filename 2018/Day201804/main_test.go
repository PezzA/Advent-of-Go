package Day201804

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	shifts, _, _ := getData(Entry.PuzzleInput())

	/*
			shifts, _, _ := getData(`[1518-11-01 00:00] Guard #10 begins shift
		[1518-11-01 00:05] falls asleep
		[1518-11-01 00:25] wakes up
		[1518-11-01 00:30] falls asleep
		[1518-11-01 00:55] wakes up
		[1518-11-01 23:58] Guard #99 begins shift
		[1518-11-02 00:40] falls asleep
		[1518-11-02 00:50] wakes up
		[1518-11-03 00:05] Guard #10 begins shift
		[1518-11-03 00:24] falls asleep
		[1518-11-03 00:29] wakes up
		[1518-11-04 00:02] Guard #99 begins shift
		[1518-11-04 00:36] falls asleep
		[1518-11-04 00:46] wakes up
		[1518-11-05 00:03] Guard #99 begins shift
		[1518-11-05 00:45] falls asleep
		[1518-11-05 00:55] wakes up`)
	*/
	//sleepyGuard := getSleepyGuard(shifts)

	//fmt.Println(sleepyGuard)
	//minsMostAsleep := getMinSleepyGuardMostAsleep(sleepyGuard, shifts)

	//fmt.Println(minsMostAsleep)

	fmt.Println(getSleepyMinute(shifts))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
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
