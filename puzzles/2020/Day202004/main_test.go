package Day202004

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	data := getData(Entry.PuzzleInput())

	test := data[0]
	Expect(test["eyr"]).Should(Equal("2029"))
	Expect(test["pid"]).Should(Equal("148714704"))

	test = data[1]
	Expect(test["byr"]).Should(Equal("2013"))
	Expect(test["ecl"]).Should(Equal("#4f9a1c"))
}

func Test_ValidPassport(t *testing.T) {
	RegisterTestingT(t)
	tests := getData(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)

	Expect(tests[0].isValid()).Should(Equal(true))
	Expect(tests[1].isValid()).Should(Equal(false))
	Expect(tests[2].isValid()).Should(Equal(true))
	Expect(tests[3].isValid()).Should(Equal(false))
}

func Test_ValidExtendedValidation(t *testing.T) {
	RegisterTestingT(t)
	data := getData(`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`)

	for _, pp := range data {
		Expect(pp.isValid()).Should(Equal(true))
		Expect(pp.passesExtendedValidation()).Should(Equal(true))
	}
}

func Test_InValidExtendedValidation(t *testing.T) {
	RegisterTestingT(t)
	data := getData(`eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`)

	for _, pp := range data {
		Expect(pp.isValid()).Should(Equal(true))
		Expect(pp.passesExtendedValidation()).Should(Equal(false))
	}
}

func Test_ValidateHeight(t *testing.T) {
	RegisterTestingT(t)
	Expect(validateHeight("193cm")).Should(Equal(true))
	Expect(validateHeight("194cm")).Should(Equal(false))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	data := getData(Entry.PuzzleInput())
	valid := 0

	for _, pp := range data {
		if pp.isValid() {
			valid++
		}
	}

	fmt.Println(valid)
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	data := getData(Entry.PuzzleInput())
	valid := 0

	for _, pp := range data {
		if pp.isValid() && pp.passesExtendedValidation() {
			valid++
		}
	}

	fmt.Println(valid)

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
