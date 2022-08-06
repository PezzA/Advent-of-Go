package Day202007

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	fmt.Println(data)
}

func Test_GetTrees(t *testing.T) {
	data := getData(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	l := findLeafs(data)

	r := getTrees(l, data)

	fmt.Println(r)

}

func Test_GetLeafs(t *testing.T) {
	data := getData(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	leafs := findLeafs(data)

	for _, l := range leafs {
		fmt.Println(l)
	}
}

func Test_GetTree(t *testing.T) {
	data := getData(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	bag := getTree("dark orange", data, 0)

	Expect(bag.colour).Should(Equal("dark orange"))
}

func Test_GetLeafsData(t *testing.T) {
	data := getData(Entry.PuzzleInput())

	leafs := findLeafs(data)

	for _, l := range leafs {
		fmt.Println(l)
	}
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

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

	//275, 274 too high
	fmt.Println(len(bags))

}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	sg := getTree("shiny gold", data, 1)

	fmt.Println(walkAdd(sg) - 1)

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
