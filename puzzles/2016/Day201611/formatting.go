package Day201611

import "fmt"

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
