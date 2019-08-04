package Day201824

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func varyingInitativePuzzleInput() string {
	return `Immune System:
4445 units each with 10125 hit points (immune to radiation) with an attack that does 20 cold damage at initiative 16
722 units each with 9484 hit points with an attack that does 130 bludgeoning damage at initiative 6
1767 units each with 5757 hit points (weak to fire, radiation) with an attack that does 27 radiation damage at initiative 4
1472 units each with 7155 hit points (weak to slashing, bludgeoning) with an attack that does 42 radiation damage at initiative 20
2610 units each with 5083 hit points (weak to slashing, fire) with an attack that does 14 fire damage at initiative 17
442 units each with 1918 hit points with an attack that does 35 fire damage at initiative 8
2593 units each with 1755 hit points (immune to bludgeoning, radiation, fire) with an attack that does 6 slashing damage at initiative 1
2593 units each with 1755 hit points (immune to bludgeoning, radiation, fire) with an attack that does 6 slashing damage at initiative 13
2593 units each with 1755 hit points (immune to bludgeoning, radiation, fire) with an attack that does 6 slashing damage at initiative 12
2593 units each with 1755 hit points (immune to bludgeoning, radiation, fire) with an attack that does 6 slashing damage at initiative 167
6111 units each with 1395 hit points (weak to bludgeoning; immune to radiation, fire) with an attack that does 1 slashing damage at initiative 14
231 units each with 3038 hit points (immune to radiation) with an attack that does 128 cold damage at initiative 15
3091 units each with 6684 hit points (weak to radiation; immune to slashing) with an attack that does 17 cold damage at initiative 19

Infection:
1929 units each with 13168 hit points (weak to bludgeoning) with an attack that does 13 fire damage at initiative 7
2143 units each with 14262 hit points (immune to radiation) with an attack that does 12 fire damage at initiative 10
2143 units each with 14262 hit points (immune to radiation) with an attack that does 12 fire damage at initiative 18
1380 units each with 20450 hit points (weak to slashing, radiation; immune to bludgeoning, fire) with an attack that does 28 cold damage at initiative 12
4914 units each with 6963 hit points (weak to slashing; immune to fire) with an attack that does 2 cold damage at initiative 11
1481 units each with 14192 hit points (weak to slashing, fire; immune to radiation) with an attack that does 17 bludgeoning damage at initiative 3
58 units each with 40282 hit points (weak to cold, slashing) with an attack that does 1346 radiation damage at initiative 9
2268 units each with 30049 hit points (immune to cold, slashing, radiation) with an attack that does 24 radiation damage at initiative 5
3562 units each with 22067 hit points with an attack that does 9 fire damage at initiative 18
4874 units each with 37620 hit points (immune to bludgeoning; weak to cold) with an attack that does 13 bludgeoning damage at initiative 1
4378 units each with 32200 hit points (weak to cold) with an attack that does 10 bludgeoning damage at initiative 2`
}

func testInput() string {
	return `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`
}

func Test_GetData(t *testing.T) {
	RegisterTestingT(t)

	immunes, infections := getData(Entry.PuzzleInput())

	Expect(immunes[0]).Should(Equal(group{"immune", 1, 4445, 10125, 20, 16, "cold", []string{}, []string{"radiation"}}))
	Expect(immunes[2]).Should(Equal(group{"immune", 3, 1767, 5757, 27, 4, "radiation", []string{"fire", "radiation"}, []string{}}))
	Expect(immunes[9]).Should(Equal(group{"immune", 10, 3091, 6684, 17, 19, "cold", []string{"radiation"}, []string{"slashing"}}))

	Expect(infections[0]).Should(Equal(group{"infection", 13, 1929, 13168, 13, 7, "fire", []string{"bludgeoning"}, []string{}}))
}

func Test_GetTargetOrder(t *testing.T) {
	RegisterTestingT(t)
	immunes, infections := getData(varyingInitativePuzzleInput())

	ol := getTargetOrder(immunes, infections)

	for _, g := range ol {
		fmt.Println(g.effectivePower(), g.initative, g)
	}
}

func Test_GetInitiativeOrder(t *testing.T) {
	RegisterTestingT(t)
	immunes, infections := getData(testInput())
	io := getInitativeOrder(immunes, infections)

	for _, g := range io {
		fmt.Println(g.effectivePower(), g.initative, g)
	}
}

func Test_AttackingPower(t *testing.T) {
	RegisterTestingT(t)

	immunes, infections := getData(testInput())

	Expect(immunes[0].units).Should(Equal(17))
	Expect(infections[0].units).Should(Equal(801))

	Expect(infections[0].attackingPower(immunes[0])).Should(Equal(185832))
	Expect(infections[0].attackingPower(immunes[1])).Should(Equal(185832))
	Expect(infections[1].attackingPower(immunes[1])).Should(Equal(107640))

	Expect(immunes[0].attackingPower(infections[0])).Should(Equal(76619))
	Expect(immunes[0].attackingPower(infections[1])).Should(Equal(153238))
	Expect(immunes[1].attackingPower(infections[0])).Should(Equal(24725))
}

func Test_RunTest(t *testing.T) {
	RegisterTestingT(t)

	immunes, infections := getData(testInput())

	battles := resolveAttackOrder(immunes, infections)

	for _, battle := range battles {
		fmt.Println(battle)

	}

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
