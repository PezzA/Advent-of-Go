package Day202004

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var hclValidator = regexp.MustCompile(`(?i)^#([0-9a-fA-F]){6}$`)
var pidValidator = regexp.MustCompile(`(?i)^[0-9]{9}$`)
var eclValidator = regexp.MustCompile(`(?i)^amb|blu|brn|gry|grn|hzl|oth$`)
var heightValidator = regexp.MustCompile(`(?i)^(?P<num>[0-9]*)(?P<type>cm|in)$`)

type passport map[string]string

// assumes data is always valid
func getData(input string) []passport {
	data := make([]passport, 0)

	current := make(passport, 0)
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			//commit the passport and create a new one
			data = append(data, current)
			current = make(passport, 0)
			continue
		}

		for _, part := range strings.Fields(line) {
			bits := strings.Split(part, ":")
			current[bits[0]] = bits[1]
		}
	}
	data = append(data, current)
	return data
}

func validNumberRange(input string, min int, max int) bool {
	val, err := strconv.Atoi(input)

	if err != nil {
		return false
	}

	return val >= min && val <= max
}

func (p passport) hasField(input string) bool {
	_, ok := p[input]
	return ok
}

func (p passport) isValid() bool {
	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range validFields {
		if !p.hasField(field) {
			return false
		}
	}

	return true
}

func validateHeight(input string) bool {
	if !heightValidator.MatchString(input) {
		return false
	}

	res := heightValidator.FindStringSubmatch(input)

	if res[2] == "in" {
		if !validNumberRange(res[1], 59, 76) {
			return false
		}
	}

	if res[2] == "cm" {
		if !validNumberRange(res[1], 150, 193) {
			return false
		}
	}
	return true
}

func (p passport) passesExtendedValidation() bool {

	if !validNumberRange(p["byr"], 1920, 2002) {
		return false
	}
	if !validNumberRange(p["iyr"], 2010, 2020) {
		return false
	}
	if !validNumberRange(p["eyr"], 2020, 2030) {
		return false
	}

	if !hclValidator.MatchString(p["hcl"]) {
		return false
	}

	if !eclValidator.MatchString(p["ecl"]) {
		return false
	}

	if !pidValidator.MatchString(p["pid"]) {
		return false
	}

	if !validateHeight(p["hgt"]) {
		return false
	}

	return true
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	valid := 0

	for _, pp := range data {
		if pp.isValid() {
			valid++
		}
	}

	return fmt.Sprintf("%v", valid)
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	data := getData(inputData)
	valid := 0

	for _, pp := range data {
		if pp.isValid() && pp.passesExtendedValidation() {
			valid++
		}
	}
	return fmt.Sprintf("%v", valid)
}
