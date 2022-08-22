package Day201611

import (
	"fmt"
	"sort"
)

func (f floor) sort() {
	f = componentSort(f)
}

func (f floor) getComponentCombos() [][]component {
	comboList := make([][]component, 0)
	// add individual components
	for _, comp := range f {
		comboList = append(comboList, []component{comp})
	}

	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f); j++ {
			if i < j {
				comboList = append(comboList, []component{f[i], f[j]})
			}
		}
	}

	return comboList
}

func (f floor) hasComponent(searchComponent component) bool {
	for _, floorComp := range f {
		if floorComp == searchComponent {
			return true
		}
	}

	return false
}

func (f floor) hasOtherGenerator(element string) bool {
	for _, component := range f {
		if !component.IsChip && component.Element != element {
			return true
		}
	}

	return false
}

func (f floor) hasGenerator(element string) bool {
	for _, component := range f {
		if !component.IsChip && component.Element == element {
			return true
		}
	}

	return false
}

func (f floor) isValid() bool {
	for _, component := range f {
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
	newFloor := make(floor, len(f))

	for i := range f {
		newFloor[i] = f[i]
	}

	return newFloor
}

func (f facility) sort() {
	for i := 0; i < len(f.Floors); i++ {
		f.Floors[i].sort()
	}
}

func (f facility) getAllComponents() []component {
	componentList := make([]component, 0)

	for _, floor := range f.Floors {
		componentList = append(componentList, floor...)
	}

	return componentSort(componentList)
}

func (f facility) canLiftGoUp() bool {
	return f.Lift < (len(f.Floors) - 1)
}

func (f facility) canLiftGoDown() bool {
	return f.Lift > 0
}

func (f facility) equals(compare facility) bool {
	if f.Lift != compare.Lift {
		return false
	}

	if len(f.Floors) != len(compare.Floors) {
		return false
	}

	for i := 0; i < len(f.Floors); i++ {
		if len(f.Floors[i]) != len(compare.Floors[i]) {
			return false
		}

		for j := 0; j < len(f.Floors[i]); j++ {
			if f.Floors[i][j] != compare.Floors[i][j] {
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

func (f floor) removeComponents(compList floor) floor {
	newList := make(floor, 0)

	for _, comp := range f {
		if !compList.hasComponent(comp) {
			newList = append(newList, comp)
		}
	}

	return newList
}

// applyMove will return a new facility based on the recieving facility
// and the component specified.
// modifer contains the direction of the elevator 1 is up , -1 is down
func (f facility) applyMove(modifier int, liftContents []component) facility {

	// create a copy of the facility to apply changes to
	newFacility := f.deepClone()
	// move elevator
	newFacility.Lift += modifier

	oldlevel, newLevel := f.Lift, newFacility.Lift

	// remove items from previous floor
	newFacility.Floors[oldlevel] =
		newFacility.Floors[oldlevel].removeComponents(liftContents)

	// add them to the target floor
	newFacility.Floors[newLevel] =
		append(newFacility.Floors[newLevel], liftContents...)

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
	return len(f.Floors[3]) == len(f.getAllComponents())
}

// maxDepth should be how many moves it has taken to solve the problem.
func (f facility) processMoves(level int, states []facility) int {
	states = append(states, f)

	fmt.Println("Processing level:", level)
	f.outputFacility()

	combos := f.Floors[f.Lift].getComponentCombos()

	if f.canLiftGoUp() {
		for _, combo := range combos {
			newFac := f.applyMove(1, combo)

			if newFac.isComplete() {
				return level
			}

			if newFac.isValid() && !newFac.alreadySeen(states) {
				fmt.Println("moving up", combo)
				level = newFac.processMoves(level+1, states)
			}
		}
	}

	if f.canLiftGoDown() {
		for _, combo := range combos {
			newFac := f.applyMove(-1, combo)
			if newFac.isComplete() {
				return level
			}

			if newFac.isValid() && !newFac.alreadySeen(states) {
				fmt.Println("moving down", combo)

				level = newFac.processMoves(level+1, states)
			}
		}
	}

	return level
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
			floor{{"Promethium", false}, {"Promethium", true}},
			floor{{"Cobalt", false}, {"Curium", false}, {"Ruthium", false}, {"Plutonium", false}},
			floor{{"Cobalt", true}, {"Curium", true}, {"Ruthium", true}, {"Plutonium", true}},
			floor{},
		},
	}
	return fmt.Sprintf("%v", inputFac.processMoves(0, []facility{})+1)
}

// PartTwo returns the solution to day11 part two
func (d dayEntry) PartTwo(input string, updateChan chan []string) string {
	return ""
}
