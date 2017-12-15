package Day201611

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

}

func Benchmark_PartOne(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartOne(data)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	data := Entry.PuzzleInput()
	for n := 0; n < b.N; n++ {
		Entry.PartTwo(data)
	}
}

/*
func Test_Day11(t *testing.T) {
	Convey("Day 11 should be able to ", t, func() {
		f := getStartingFacility(Entry.PuzzleInput())

		hydrogenMicrochip := component{0, "hydrogen", true}
		lithiumMicrochip := component{0, "lithium", true}
		hydrogenGenerator := component{1, "hydrogen", false}
		lithiumGenerator := component{2, "lithium", false}

		testComponentList := make(componentList, 0)
		testComponentList = append(testComponentList, hydrogenMicrochip)
		testComponentList = append(testComponentList, lithiumMicrochip)
		testComponentList = append(testComponentList, hydrogenGenerator)
		testComponentList = append(testComponentList, lithiumGenerator)

		Convey("-> Parse the puzzle input", func() {
			So(f, ShouldResemble, facility{0, testComponentList})
		})

		Convey("-> Draw the facility", func() {
			f.drawFacility()
		})

		Convey("-> Should be able to tell if 2 components are the same", func() {
			So(lithiumGenerator.equals(lithiumMicrochip), ShouldBeFalse)
			So(lithiumGenerator.equals(component{2, "lithium", false}), ShouldBeTrue)
		})

		Convey("-> Should be able to sort a component list", func() {
			tempL := testComponentList

			sort.Sort(tempL)
			So(tempL[0], ShouldResemble, hydrogenMicrochip)
			So(tempL[1], ShouldResemble, hydrogenGenerator)
			So(tempL[2], ShouldResemble, lithiumMicrochip)
			So(tempL[3], ShouldResemble, lithiumGenerator)
		})

		Convey("-> Get all the stuff on the specified floor", func() {
			So(f.getItemsOnFloor(0)[0], ShouldResemble, hydrogenMicrochip)
			So(f.getItemsOnFloor(0)[1], ShouldResemble, lithiumMicrochip)
			So(f.getItemsOnFloor(1)[0], ShouldResemble, hydrogenGenerator)
			So(f.getItemsOnFloor(2)[0], ShouldResemble, lithiumGenerator)
			So(len(f.getItemsOnFloor(3)), ShouldResemble, 0)
		})

		Convey("-> Give all the cominbinations for the stuff on the floor", func() {
			items := elevatorCombinations(f.getItemsOnFloor(0))
			So(items[0], ShouldResemble, []component{hydrogenMicrochip})
			So(items[1], ShouldResemble, []component{lithiumMicrochip})
			So(items[2], ShouldResemble, []component{hydrogenMicrochip, lithiumMicrochip})
		})

		Convey("-> Should be able to tell 2 facilties have the same state", func() {
			x := getStartingFacility(Entry.PuzzleInput())

			y := getStartingFacility(Entry.PuzzleInput())

			So(x.equals(y), ShouldBeTrue)

			y.elevator = 2

			So(x.equals(y), ShouldBeFalse)

			y.elevator = 0
			y.components[0].floor = 3

			So(x.equals(y), ShouldBeFalse)

			y.components[0].floor = 0

			So(x.equals(y), ShouldBeTrue)

		})

		Convey("-> Should be able to detect if moves are valid", func() {

			log.Println("############################################################")
			solve(f, make([]facility, 0))

		})
	})
}
*/
