package Day201823

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDay201823(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day201815 Suite")
}

var _ = Describe("Point", func() {
	It("Should be able to load the data", func(){

		bots:=getData(Entry.PuzzleInput())
		Expect(bots[0]).Should(Equal(nanobot{point3D{-3079732, 26664397, 4300940}, 64337129}))
	})

	It ("Should be able to work out the distance between 2 points", func(){
		origin := point3D{0, 0, 0}
		Expect(origin.getMDistance(point3D{0, 0, 0})).Should(Equal(0))
		Expect(origin.getMDistance(point3D{1, 0, 0})).Should(Equal(1))
		Expect(origin.getMDistance(point3D{4, 0, 0})).Should(Equal(4))
		Expect(origin.getMDistance(point3D{0, 2, 0})).Should(Equal(2))
		Expect(origin.getMDistance(point3D{0, 5, 0})).Should(Equal(5))
		Expect(origin.getMDistance(point3D{0, 0, 3})).Should(Equal(3))
		Expect(origin.getMDistance(point3D{1, 1, 1})).Should(Equal(3))
		Expect(origin.getMDistance(point3D{1, 1, 2})).Should(Equal(4))
		Expect(origin.getMDistance(point3D{1, 3, 1})).Should(Equal(5))
	})

	It("Should be able to find the nanobot with the highest radions", func() {
		bots:=getData(`pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1`)

		Expect(getHighestRadius(bots)).Should(Equal(nanobot{point3D{0, 0, 0}, 4}))
	})

	It("Should be able to detect the number of bots within range of a bot", func() {
		bots:=getData(`pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1`)

		rBot := getHighestRadius(bots)
		botsInRange := rBot.getBotsInRange(bots)

		Expect(len(botsInRange)).Should(Equal(7))
	})
})