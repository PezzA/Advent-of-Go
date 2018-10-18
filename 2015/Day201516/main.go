package Day201516

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry holds wraps the data and runner interfaces for this puzzle
var Entry dayEntry

type dayEntry bool

type aunt struct {
	id int
	properties map[string]int
}

func getData(input string) []aunt{
	aunties := make([]aunt, 0)

	for _, line := range strings.Split(input, "\n"){
		bits := strings.Fields(line)

		id, _ := strconv.Atoi(strings.Replace(bits[1],":","",-1))

		prop1  := strings.Replace(bits[2], ":","",-1)
		val1, _  := strconv.Atoi(strings.Replace(bits[3], ",","",-1))
		prop2  := strings.Replace(bits[4], ":","",-1)
		val2, _  := strconv.Atoi(strings.Replace(bits[5], ",","",-1))
		prop3  := strings.Replace(bits[6], ":","",-1)
		val3, _  := strconv.Atoi(strings.Replace(bits[7], ",","",-1))

		auntie := aunt{id, make(map[string]int)}
		auntie.properties[prop1] = val1
		auntie.properties[prop2] = val2
		auntie.properties[prop3] = val3
		aunties = append(aunties, auntie)
	}

	return aunties
}

func readTickerTape() map[string]int {
	tickerMap := make(map[string]int)

	tickerMap["children"] = 3
	tickerMap["cats"] = 7
	tickerMap["samoyeds"] = 2
	tickerMap["pomeranians"] = 3
	tickerMap["akitas"] = 0
	tickerMap["vizslas"] = 0
	tickerMap["goldfish"] = 5
	tickerMap["trees"] = 3
	tickerMap["cars"] = 2
	tickerMap["perfumes"] = 1

	return tickerMap
}

func (td dayEntry) Describe() (int, int, string) {
	return 2015, 16, "Aunt Sue"
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	aunties := getData(inputData)
	tickerTape := readTickerTape()
	targetAunt :=0
	for _, aunt := range aunties{
		hits := 0

		for k, v := range aunt.properties{
			if v == tickerTape[k]{
				hits++
			}
		}

		if hits == 3 {
			targetAunt = aunt.id
			break
		}
	}

	return fmt.Sprintf("%v", targetAunt)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	aunties := getData(inputData)
	tickerTape := readTickerTape()
	targetAunt :=0
	for _, aunt := range aunties{
		hits := 0

		for k, v := range aunt.properties{

			if k == "cats" || k == "trees" {
				if v > tickerTape[k]{
					hits++
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if v < tickerTape[k]{
					hits++
				}
			} else if v == tickerTape[k]{
				hits++
			}
		}

		if hits == 3 {
			targetAunt = aunt.id
			break
		}
	}

	return fmt.Sprintf("%v", targetAunt)
}
