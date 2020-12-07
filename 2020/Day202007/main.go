package Day202007

import (
	"fmt"
	"strconv"
	"strings"
)

type content struct {
	colour   string
	quantity int
}

type instruction struct {
	colour   string
	contents []content
}

func getData(input string) []instruction {
	data := make([]instruction, 0)

	for _, line := range strings.Split(input, "\n") {
		bits := strings.Split(line, "bags contain")
		bags := strings.Split(strings.TrimSpace(bits[1]), ",")

		contents := make([]content, 0)
		for _, b := range bags {
			if b == "no other bags" {
				break
			}

			bagBits := strings.Fields(b)

			quantity, _ := strconv.Atoi(bagBits[0])

			contents = append(contents, content{
				quantity: quantity,
				colour:   fmt.Sprintf("%s %s", bagBits[1], bagBits[2]),
			})
		}
		data = append(data, instruction{
			colour:   strings.TrimSpace(bits[0]),
			contents: contents,
		})
	}

	return data
}

type bag struct {
	colour   string
	quantity int
	contents []*bag
}

func getInstruction(colour string, ins []instruction) (bool, instruction) {
	for _, i := range ins {
		if i.colour == colour {
			return true, i
		}
	}
	return false, instruction{}
}

func findLeafs(ins []instruction) []string {
	leafs := make([]string, 0)

	for _, i := range ins {

		checkColor := i.colour
		found := false
		for _, ic := range ins {
			for _, icc := range ic.contents {
				if icc.colour == checkColor {
					found = true
				}
			}
		}

		if !found {
			leafs = append(leafs, checkColor)
		}

	}

	return leafs
}

func getTrees(roots []string, ins []instruction) []*bag {
	data := make([]*bag, len(roots))

	for i := range roots {
		data[i] = getTree(roots[i], ins, 0)
	}

	return data
}

func getTree(color string, ins []instruction, q int) *bag {
	_, rootIns := getInstruction(color, ins)

	rootBag := &bag{
		colour:   rootIns.colour,
		contents: make([]*bag, 0),
		quantity: q,
	}

	for _, c := range rootIns.contents {
		rootBag.contents = append(rootBag.contents, getTree(c.colour, ins, c.quantity))
	}

	return rootBag
}

var seen = make([]string, 0)

func walkAdd(b *bag) int {

	contents := b.quantity
	for _, sb := range b.contents {
		contents += walkAdd(sb) * b.quantity

	}
	return contents
}

// assumption, shiny gold bags wont eventually contain other shiny gold bags
//
func walkBag(b *bag, opened []string) {
	if b.colour == "shiny gold" {
		seen = append(seen, opened...)

	}

	opened = append(opened, b.colour)

	for _, sb := range b.contents {
		walkBag(sb, opened)
	}

}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)

	leafs := findLeafs(data)

	seen = make([]string, 0)

	for _, l := range leafs {
		t := getTree(l, data, 0)

		walkBag(t, []string{})
	}

	bags := make(map[string]int, 0)

	for _, s := range seen {
		if _, ok := bags[s]; ok {
			bags[s]++
		} else {
			bags[s] = 1
		}
	}

	return fmt.Sprintf("%v", len(bags))
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	sg := getTree("shiny gold", data, 1)
	return fmt.Sprintf("%v", walkAdd(sg)-1)
}
