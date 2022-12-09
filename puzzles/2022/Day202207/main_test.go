package Day202207

import (
	"math"
	"testing"

	. "github.com/onsi/gomega"
)

var sampleInput string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_ReadData(t *testing.T) {
	RegisterTestingT(t)

	fs := getData(sampleInput)

	Expect(len(fs.dirs)).Should(Equal(2))
	Expect(len(fs.files)).Should(Equal(2))
}

func Test_WalkTree(t *testing.T) {
	RegisterTestingT(t)

	fs := getData(sampleInput)

	calcFolderSize(fs)

	Expect(fs.size).Should(Equal(48381165))

	size := getFoldersUnderSizeLimit(fs, 100000)

	Expect(size).Should(Equal(95437))
}

func Test_Smallest(t *testing.T) {

	RegisterTestingT(t)

	fs := getData(sampleInput)

	totalSize := calcFolderSize(fs)
	maxSize, spaceNeeded := 70000000, 30000000
	// 1735494
	sizeToReclaim := spaceNeeded - (maxSize - totalSize)

	smallestReclaimDirSize := getSmallestToReclaim(fs, sizeToReclaim, math.MaxInt32)

	Expect(smallestReclaimDirSize).Should(Equal(24933642))
}

func Test_PartOne(t *testing.T) {
	RegisterTestingT(t)
	Expect(Entry.PartOne(Entry.PuzzleInput(), nil)).Should(Equal("1648397"))
}

func Test_PartTwo(t *testing.T) {
	RegisterTestingT(t)
	Expect(Entry.PartTwo(Entry.PuzzleInput(), nil)).Should(Equal("1815525"))
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
