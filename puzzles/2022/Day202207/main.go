package Day202207

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type directory struct {
	dirs   map[string]*directory
	files  map[string]int
	parent *directory
	size   int // size in bytes of the directory
}

func newDir(parent *directory) *directory {
	blankDirs := make(map[string]*directory, 0)
	files := make(map[string]int)
	return &directory{
		dirs:   blankDirs,
		files:  files,
		parent: parent,
		size:   0,
	}
}

func getData(input string) *directory {
	lines := strings.Split(input, "\n")

	root := newDir(nil)

	currDir := root
	for _, line := range lines {
		bits := strings.Split(line, " ")

		switch bits[0] {
		case "$":
			switch bits[1] {
			case "cd":
				switch bits[2] {
				case "..":
					currDir.parent.size += currDir.size
					currDir = currDir.parent
				case "/":
					currDir = root
				default:
					currDir = currDir.dirs[bits[2]]
				}
			case "ls":
				continue
			}
		case "dir":
			currDir.dirs[bits[1]] = newDir(currDir)
		default:
			size, _ := strconv.Atoi(bits[0])
			currDir.files[bits[1]] = size
			currDir.size += size
		}
	}

	return root
}

func calcFolderSize(fs *directory) int {
	size := 0

	for _, fileSize := range fs.files {
		size += fileSize
	}

	for _, dir := range fs.dirs {
		size += calcFolderSize(dir)
	}

	fs.size = size

	return size
}

func getFoldersUnderSizeLimit(fs *directory, limit int) int {
	count := 0

	if fs.size < limit {
		count += fs.size
	}

	for _, dir := range fs.dirs {
		count += getFoldersUnderSizeLimit(dir, limit)
	}

	return count
}

func (td dayEntry) PartOne(inputData string, updateChan chan []string) string {
	fs := getData(inputData)
	calcFolderSize(fs)

	return fmt.Sprint(getFoldersUnderSizeLimit(fs, 100000))
}

func getSmallestToReclaim(fs *directory, reclaimThreshold int, currentSmallest int) int {

	if fs.size >= reclaimThreshold && fs.size < currentSmallest {
		currentSmallest = fs.size
	}

	for _, dir := range fs.dirs {
		dirSize := getSmallestToReclaim(dir, reclaimThreshold, currentSmallest)

		if dirSize < currentSmallest {
			currentSmallest = dirSize
		}

	}
	return currentSmallest
}

func (td dayEntry) PartTwo(inputData string, updateChan chan []string) string {
	fs := getData(inputData)
	totalSize := calcFolderSize(fs)
	maxSize, spaceNeeded := 70000000, 30000000
	// 1735494
	sizeToReclaim := spaceNeeded - (maxSize - totalSize)

	smallestReclaimDirSize := getSmallestToReclaim(fs, sizeToReclaim, math.MaxInt32)
	return fmt.Sprint(smallestReclaimDirSize)
}
