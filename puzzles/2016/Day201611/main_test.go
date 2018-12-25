package Day201611

import (
	"testing"

	"fmt"

	. "github.com/onsi/gomega"
)

func puzzleTestData() ([]string, []string) {
	return []string{`The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
	The second floor contains a hydrogen generator.
	The third floor contains a lithium generator.
	The fourth floor contains nothing relevant.`}, []string{}
}

func getTestFacility() facility {
	return facility{
		Lift: 0,
		Floors: []floor{
			{[]component{{"Hydrogen", true}, {"Lithium", true}}},
			{[]component{{"Hydrogen", false}}},
			{[]component{{"Lithium", false}}},
			{[]component{}},
		},
	}
}

func Test_facilityClone(t *testing.T) {
	RegisterTestingT(t)

	origFloor := floor{[]component{{"Lithium", false}}}
	clonedFloor := origFloor.deepClone()

	Expect(origFloor).Should(Equal(clonedFloor))

	clonedFloor.Components = append(clonedFloor.Components, component{"Saphirite", false})

	Expect(origFloor).ShouldNot(Equal(clonedFloor))

	newFacility := getTestFacility()

	clonedFacility := newFacility.deepClone()

	Expect(newFacility).Should(Equal(clonedFacility))
}

func Test_Equality(t *testing.T) {
	RegisterTestingT(t)
	fac1 := getTestFacility()
	fac2 := getTestFacility()

	Expect(fac1.equals(fac2)).Should(Equal(true))
}
func Test_floorEquals(t *testing.T) {
	RegisterTestingT(t)

	// both list are in different incorrect orders, both should be the same at the end
	floor1 := floor{[]component{{"Hydrogen", false}, {"Hydrogen", true}, {"Lithium", true}}}
	floor2 := floor{[]component{{"Lithium", true}, {"Hydrogen", true}, {"Hydrogen", false}}}

	Expect(floor1).ShouldNot(Equal(floor2))

	floor1.sort()
	floor2.sort()

	Expect(floor1).Should(Equal(floor2))
}

func Test_getComponentCombos(t *testing.T) {
	RegisterTestingT(t)
	floor1 := floor{[]component{{"Hydrogen", false}, {"Hydrogen", true}, {"Lithium", true}}}

	comboList := floor1.getComponentCombos()
	fmt.Println(comboList)
	Expect(len(comboList)).Should(Equal(6))
}

func Test_drawFacility(t *testing.T) {
	getTestFacility().outputFacility()
}

func Test_shortCode(t *testing.T) {
	RegisterTestingT(t)
	// just testing the scenario's I'll hit, don't care about edges
	testComponent := component{"Uranium", true}
	Expect(testComponent.shortCode()).Should(Equal("UM"))

	testComponent2 := component{"Crontinium", false}
	Expect(testComponent2.shortCode()).Should(Equal("CG"))
}

func Test_floorIsValid(t *testing.T) {
	RegisterTestingT(t)

	testFloor := floor{[]component{
		{"Hydrogen", true},
		{"Lithium", true}}}

	Expect(testFloor.isValid()).Should(Equal(true))

	testFloor2 := floor{[]component{
		{"Hydrogen", true},
		{"Lithium", true},
		{"Lithium", false},
		{"Hydrogen", false}}}

	Expect(testFloor2.isValid()).Should(Equal(true))

	testFloor3 := floor{[]component{
		{"Hydrogen", true},
		{"Lithium", true},
		{"Lithium", false}}}

	Expect(testFloor3.isValid()).Should(Equal(false))
}

func Test_applyMove(t *testing.T) {
	startFacility := getTestFacility()

	fmt.Println("(1)")
	startFacility.outputFacility()

	updatedFacility := startFacility.applyMove(1, []component{{"Hydrogen", true}})
	fmt.Println("")
	fmt.Println("(2)")
	updatedFacility.outputFacility()
}

func Test_processMoves(t *testing.T) {
	startFacility := getTestFacility()

	val := startFacility.processMoves(0, []facility{}, 0, nil)

	fmt.Println(val + 1)
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	inputFac := facility{
		Lift: 0,
		Floors: []floor{
			{[]component{{"Promethium", false}, {"Promethium", true}}},
			{[]component{{"Cobalt", false}, {"Curium", false}, {"Ruthium", false}, {"Plutonium", false}}},
			{[]component{{"Cobalt", true}, {"Curium", true}, {"Ruthium", true}, {"Plutonium", true}}},
			{[]component{}},
		},
	}
	fmt.Println(inputFac.processMoves(0, []facility{}, 0, nil) + 1)
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
