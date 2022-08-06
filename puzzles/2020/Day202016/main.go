package Day202016

import (
	"fmt"
	"strconv"
	"strings"
)

type ticketField struct {
	name   string
	r1Low  int
	r1High int
	r2Low  int
	r2High int
}

type ticket []int

func getData(input string) ([]ticketField, ticket, []ticket) {
	lines := strings.Split(input, "\n")

	ticketFields := make([]ticketField, 0)

	myTicket := make(ticket, 0)
	nearbyTickets := make([]ticket, 0)

	gettingMyTicket, gettingNearbyTickets := false, false
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line == "your ticket:" && !gettingMyTicket {
			gettingMyTicket = true
			continue
		}

		if line == "nearby tickets:" && gettingMyTicket {
			gettingNearbyTickets = true
			continue
		}

		if !gettingMyTicket && !gettingNearbyTickets {
			bits := strings.Split(line, ":")
			rangeBits := strings.Split(strings.TrimSpace(bits[1]), " or ")
			range1bits := strings.Split(rangeBits[0], "-")
			range2bits := strings.Split(rangeBits[1], "-")

			r1l, _ := strconv.Atoi(range1bits[0])
			r1h, _ := strconv.Atoi(range1bits[1])
			r2l, _ := strconv.Atoi(range2bits[0])
			r2h, _ := strconv.Atoi(range2bits[1])

			ticketFields = append(ticketFields, ticketField{
				name:   bits[0],
				r1Low:  r1l,
				r1High: r1h,
				r2Low:  r2l,
				r2High: r2h,
			})
			continue
		}

		if gettingMyTicket && !gettingNearbyTickets {
			fields := strings.Split(line, ",")

			for _, v := range fields {
				intV, _ := strconv.Atoi(v)

				myTicket = append(myTicket, intV)
			}
			continue
		}

		if gettingNearbyTickets {
			fields := strings.Split(line, ",")
			ticket := make(ticket, 0)
			for _, v := range fields {
				intV, _ := strconv.Atoi(v)

				ticket = append(ticket, intV)
			}

			nearbyTickets = append(nearbyTickets, ticket)
		}
	}

	return ticketFields, myTicket, nearbyTickets
}

func (t ticket) invalidValues(rules []ticketField) (bool, []int) {

	invalids := make([]int, 0)

	for _, val := range t {

		validAnyField := false

		for _, rule := range rules {
			if val >= rule.r1Low && val <= rule.r1High {
				validAnyField = true
			}
			if val >= rule.r2Low && val <= rule.r2High {
				validAnyField = true
			}
		}

		if !validAnyField {
			invalids = append(invalids, val)
		}
	}

	return len(invalids) > 0, invalids
}

type match struct {
	field string
	col   int
}

func isFoundCol(index int, foundMatches []match) bool {
	for _, match := range foundMatches {
		if match.col == index {
			return true
		}
	}

	return false
}

func isFoundrule(name string, foundMatches []match) bool {
	for _, match := range foundMatches {
		if match.field == name {
			return true
		}
	}

	return false
}
func (tf ticketField) impliedField(samples []ticket, foundMatches []match) []match {

	matches := make([]match, 0)
	// each field in the sample
	for i := 0; i < len(samples[0]); i++ {
		if isFoundCol(i, foundMatches) {
			continue
		}

		fullMatch := true

		for _, samp := range samples {
			inR1, inR2 := false, false
			if samp[i] >= tf.r1Low && samp[i] <= tf.r1High {
				inR1 = true
			}
			if samp[i] >= tf.r2Low && samp[i] <= tf.r2High {
				inR2 = true
			}

			if !inR1 && !inR2 {
				fullMatch = false
				break
			}
		}

		if fullMatch {
			matches = append(matches, match{tf.name, i})
		}
	}

	return matches
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	fields, _, others := getData(inputData)

	tot := 0

	for _, tic := range others {
		_, invalids := tic.invalidValues(fields)

		for i := range invalids {
			tot += invalids[i]
		}
	}
	return fmt.Sprintf("%v", tot)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	fields, myTicket, others := getData(inputData)

	validOthers := make([]ticket, 0)

	for _, other := range others {
		invalid, _ := other.invalidValues(fields)
		if !invalid {
			validOthers = append(validOthers, other)
		}
	}

	foundMatches := make([]match, 0)

	for {
		for _, rule := range fields {

			if isFoundrule(rule.name, foundMatches) {
				continue
			}
			matches := rule.impliedField(validOthers, foundMatches)

			if len(matches) == 1 {
				// hit!
				foundMatches = append(foundMatches, matches[0])
			}
		}

		if len(foundMatches) == len(fields) {
			// they are all allocated
			break
		}
	}

	tot := 1
	for _, fm := range foundMatches {
		if strings.HasPrefix(fm.field, "departure ") {
			tot *= myTicket[fm.col]
		}
	}

	return fmt.Sprintf("%v", tot)
}
