package Day202013

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	earliest, schedules := getData(Entry.PuzzleInput())

	Expect(earliest).Should(Equal(1003681))

	Expect(schedules[0]).Should(Equal(schedule{true, 23, 0}))
	Expect(schedules[1]).Should(Equal(schedule{false, 0, 1}))
	Expect(schedules[len(schedules)-1]).Should(Equal(schedule{true, 29, len(schedules) - 1}))

}

func Test_TimeToNextBus(t *testing.T) {
	RegisterTestingT(t)

	Expect(schedule{true, 7, 0}.timeToNextBus(939)).Should(Equal(6))
	Expect(schedule{true, 13, 0}.timeToNextBus(939)).Should(Equal(10))
	Expect(schedule{true, 59, 0}.timeToNextBus(939)).Should(Equal(5))
}

func Test_GetSync(t *testing.T) {
	RegisterTestingT(t)
	/*
		s1 := schedule{true, 7, 0}
		s2 := schedule{true, 13, 1}
		s3 := schedule{true, 59, 4}
		s4 := schedule{true, 31, 6}
		s5 := schedule{true, 19, 7}

		//nextStep, _ := getStep(s1, s2, s1.timeStamp)

		Expect(nextStep).Should(Equal(91))

		nextStep, _ = getStep(s2, s3, nextStep)

		Expect(nextStep).Should(Equal(5369))

		nextStep, _ = getStep(s3, s4, nextStep)

		Expect(nextStep).Should(Equal(166439))

		nextStep, _ = getStep(s4, s5, nextStep)

		Expect(nextStep).Should(Equal(166439))
	*/
}

func Test_GetSync2(t *testing.T) {
	RegisterTestingT(t)

	s1 := schedule{true, 17, 0}
	s2 := schedule{true, 13, 2}
	s3 := schedule{true, 19, 3}

	s2offSet, s2Step := getStep(s1, s2, 0, s1.timeStamp, false)

	fmt.Println(s2offSet, s2Step)

	s3offSet, s3Step := getStep(s2, s3, s2offSet, s2Step, true)

	fmt.Println(s3offSet, s3Step)
}

func Test_Output(t *testing.T) {

	s1 := schedule{true, 17, 0}
	s2 := schedule{true, 13, 2}
	s3 := schedule{true, 19, 3}

	fmt.Printf("%-10v%-10v%-10v%-10v\n", "time", s1, s2, s3)
	for i := 1; i < 5000; i++ {

		s1T, s2T, s3T := ".", ".", "."

		if s1.isSyncWithSlot(i) {
			s1T = "D"
		}

		if s2.isSyncWithSlot(i) {
			s2T = "D"
		}

		if s3.isSyncWithSlot(i) {
			s3T = "D"
		}
		fmt.Printf("%-10v%-10v%-10v%-10v\n", i, s1T, s2T, s3T)
	}
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	offset, schedules := getData(Entry.PuzzleInput())

	waitTime := -1
	var sched = schedule{false, 0, 0}

	for i := range schedules {
		if !schedules[i].active {
			continue

		}
		swt := schedules[i].timeToNextBus(offset)

		if swt < waitTime || waitTime == -1 {
			waitTime = swt
			sched = schedules[i]
		}
	}

	Expect(sched.timeStamp * waitTime).Should(Equal(295))
}

func Test_GetSync3(t *testing.T) {
	RegisterTestingT(t)

	s1 := schedule{true, 67, 0}
	s2 := schedule{true, 7, 1}
	s3 := schedule{true, 59, 2}
	s4 := schedule{true, 61, 3}

	s2offSet, s2Step := getStep(s1, s2, 0, s1.timeStamp, false)
	s3offSet, s3Step := getStep(s2, s3, s2offSet, s2Step, false)
	s4offSet, _ := getStep(s3, s4, s3offSet, s3Step, true)

	fmt.Println(s4offSet)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	_, schedules := getData(Entry.PuzzleInput())

	activeSchedules := make([]schedule, 0)

	for i := range schedules {
		if schedules[i].active {
			activeSchedules = append(activeSchedules, schedules[i])
		}
	}

	offset, step := 0, activeSchedules[0].timeStamp

	for i := 0; i < len(activeSchedules)-2; i++ {
		if activeSchedules[i].active {
			offset, step = getStep(activeSchedules[i], activeSchedules[i+1], offset, step, false)
		}
	}

	finalOffset, _ := getStep(activeSchedules[len(activeSchedules)-2], activeSchedules[len(activeSchedules)-1], offset, step, true)

	fmt.Println(finalOffset)

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
