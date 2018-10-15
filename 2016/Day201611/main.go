package Day201611

import (
	"fmt"
	"sort"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type component struct {
	Element string
	IsChip  bool
}

type floor struct {
	Components []component
}

type facility struct {
	Lift   int
	Floors []floor
}

type stateList []facility

func (c component) shortCode() string {
	componentType := ""

	if c.IsChip {
		componentType = "M"
	} else {
		componentType = "G"
	}

	return fmt.Sprintf("%v%v", string(c.Element[0]), componentType)
}

func (c component) String() string {
	if c.IsChip {
		return fmt.Sprintf("{%v Microchip}", c.Element)
	}

	return fmt.Sprintf("{%v Generator}", c.Element)
}

func (f floor) sort() {
	f.Components = componentSort(f.Components)
}

func (f floor) getComponentCombos() [][]component {
	comboList := make([][]component, 0)
	// add individual components
	for _, comp := range f.Components {
		comboList = append(comboList, []component{comp})
	}

	for i := 0; i < len(f.Components); i++ {
		for j := 0; j < len(f.Components); j++ {
			if i < j {
				comboList = append(comboList, []component{f.Components[i], f.Components[j]})
			}
		}
	}

	return comboList
}

func (f floor) hasComponent(searchComponent component) bool {
	for _, floorComp := range f.Components {
		if floorComp == searchComponent {
			return true
		}
	}

	return false
}

func (f floor) hasOtherGenerator(element string) bool {
	for _, component := range f.Components {
		if !component.IsChip && component.Element != element {
			return true
		}
	}

	return false
}

func (f floor) hasGenerator(element string) bool {
	for _, component := range f.Components {
		if !component.IsChip && component.Element == element {
			return true
		}
	}

	return false
}

func (f floor) isValid() bool {
	for _, component := range f.Components {
		if !component.IsChip {
			continue
		}

		if !f.hasOtherGenerator(component.Element) {
			continue
		} else {
			if !f.hasGenerator(component.Element) {
				return false
			}
		}

	}

	return true
}

func (f floor) deepClone() floor {
	floorCopy := floor{make([]component, len(f.Components))}
	copy(floorCopy.Components, f.Components)
	return floorCopy
}

func (f facility) sort() {
	for i := 0; i < len(f.Floors); i++ {
		f.Floors[i].sort()
	}
}

func (f facility) getAllComponents() []component {
	componentList := make([]component, 0)

	for _, floor := range f.Floors {
		componentList = append(componentList, floor.Components...)
	}

	return componentSort(componentList)
}

func (f facility) canLiftGoUp() bool {
	return f.Lift < (len(f.Floors) - 1)
}

func (f facility) canLiftGoDown() bool {
	return f.Lift > 0
}

func (f facility) outputFacility() {
	diagram := f.drawFacility()

	for index := range diagram {
		currIndex := (len(diagram) - index) - 1
		fmt.Println(diagram[currIndex])
	}
}

func (f facility) drawFacility() []string {
	drawing := make([]string, len(f.Floors))

	allComponents := f.getAllComponents()

	for index := range drawing {
		drawing[index] += fmt.Sprintf("F%v ", index+1)

		if f.Lift == index {
			drawing[index] += "E  "
		} else {
			drawing[index] += ".  "
		}

		for _, comp := range allComponents {
			if f.Floors[index].hasComponent(comp) {
				drawing[index] += comp.shortCode() + " "
			} else {
				drawing[index] += ".  "
			}
		}
	}

	return drawing
}

func (f facility) equals(compare facility) bool {
	if f.Lift != compare.Lift {
		return false
	}

	if len(f.Floors) != len(compare.Floors) {
		return false
	}

	for i := 0; i < len(f.Floors); i++ {
		if len(f.Floors[i].Components) != len(compare.Floors[i].Components) {
			return false
		}

		for j := 0; j < len(f.Floors[i].Components); j++ {
			if f.Floors[i].Components[j] != compare.Floors[i].Components[j] {
				return false
			}
		}
	}

	return true
}

func (f facility) isValid() bool {
	for _, floor := range f.Floors {
		if !floor.isValid() {
			return false
		}
	}
	return true
}

func (f facility) deepClone() facility {
	facilityCopy := facility{}
	facilityCopy.Floors = make([]floor, 0)
	facilityCopy.Lift = f.Lift

	for _, floor := range f.Floors {
		facilityCopy.Floors = append(facilityCopy.Floors, floor.deepClone())
	}

	return facilityCopy
}

func (f facility) applyMove(modifier int, contents []component) facility {
	newFacility := f.deepClone()

	testFloor := floor{contents}
	// remove items from current floor

	newList := make([]component, 0)
	for _, floorComponent := range newFacility.Floors[newFacility.Lift].Components {
		if !testFloor.hasComponent(floorComponent) {
			newList = append(newList, floorComponent)
		}
	}

	newFacility.Floors[newFacility.Lift].Components = newList

	// move elevator
	newFacility.Lift += modifier

	// add items to floor
	for _, newComp := range contents {
		newFacility.Floors[newFacility.Lift].Components = append(newFacility.Floors[newFacility.Lift].Components, newComp)
	}

	newFacility.sort()

	return newFacility
}

func (f facility) alreadySeen(states []facility) bool {

	for _, state := range states {
		if f.equals(state) {
			return true
		}
	}
	return false

}

func (f facility) isComplete() bool {
	return len(f.Floors[3].Components) == len(f.getAllComponents())
}

func (f facility) processMoves(level int, states []facility, maxDepth int, updateChan chan []string) int {
	if maxDepth != 0 && level >= maxDepth {
		return maxDepth
	}

	states = append(states, f)

	combos := f.Floors[f.Lift].getComponentCombos()

	if f.canLiftGoUp() {
		for _, combo := range combos {
			newFac := f.applyMove(1, combo)

			if newFac.isComplete() {
				if updateChan != nil {
					updateChan <- []string{fmt.Sprintf("%v", level)}
				}
				return level
			}

			if newFac.isValid() {
				if !newFac.alreadySeen(states) {
					maxDepth = newFac.processMoves(level+1, states, maxDepth, updateChan)
				}
			}
		}
	}

	if f.canLiftGoDown() {
		for _, combo := range combos {
			newFac := f.applyMove(-1, combo)
			if newFac.isComplete() {
				if updateChan != nil {
					updateChan <- []string{fmt.Sprintf("%v", level)}
				}
				return level
			}

			if newFac.isValid() {
				if !newFac.alreadySeen(states) {
					maxDepth = newFac.processMoves(level+1, states, maxDepth, updateChan)
				}
			}
		}
	}

	return maxDepth
}

func (d dayEntry) Describe() (int, int, string) {
	return 2016, 11, "Radioisotope Thermoelectric Generators"
}

func componentSort(list []component) []component {
	sort.Slice(list, func(i, j int) bool {
		if list[i].Element != list[j].Element {
			return list[i].Element < list[j].Element
		}

		return !list[i].IsChip
	})

	return list
}

// PartOne returns the solution to day11 part one
func (d dayEntry) PartOne(input string, updateChan chan []string) string {

	inputFac := facility{
		Lift: 0,
		Floors: []floor{
			{[]component{{"Promethium", false}, {"Promethium", true}}},
			{[]component{{"Cobalt", false}, {"Curium", false}, {"Ruthium", false}, {"Plutonium", false}}},
			{[]component{{"Cobalt", true}, {"Curium", true}, {"Ruthium", true}, {"Plutonium", true}}},
			{[]component{}},
		},
	}
	return fmt.Sprintf("%v", inputFac.processMoves(0, []facility{}, 0, updateChan)+1)
}

// PartTwo returns the solution to day11 part two
func (d dayEntry) PartTwo(input string, updateChan chan []string) string {
	return ""
}
