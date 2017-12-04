package Day201611

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func (f facility) getLine(floor int) string {
	line := ""

	line += "F" + strconv.Itoa(floor+1) + " "

	if f.elevator == floor {
		line += "E  "
	} else {
		line += ".  "
	}

	sortedComponents := f.components

	sort.Sort(sortedComponents)

	for _, component := range sortedComponents {
		if component.floor == floor {
			line += component.shortName() + " "
		} else {
			line += ".  "
		}
	}

	return line
}

func (f facility) drawFacility() {
	log.Println("-----------------------------------")
	for i := 4 - 1; i >= 0; i-- {
		fmt.Println(f.getLine(i))
	}
}
