package Day201815

import "github.com/pezza/advent-of-code/puzzles/Common"

type distanceMap map[common.Point]int

func (dm distanceMap) filterAdjacent(point common.Point) distanceMap {
	retVal := make(distanceMap)

	for _, dir := range common.Cardinal4 {
		tp := point.Add(dir)

		if _, ok := dm[tp]; ok {
			retVal[tp] = dm[tp]
		}
	}
	return retVal
}

func (dm distanceMap) getTop() common.Point {
	minPoint := -1
	var targetLocation common.Point

	for k, v := range dm {

		// if its the same value, use it if it is higher up in reading order
		if minPoint == v {
			if k.Y < targetLocation.Y {
				minPoint = v
				targetLocation = k
			} else if k.Y == targetLocation.Y {
				if k.X < targetLocation.X {
					minPoint = v
					targetLocation = k
				}
			}
		}

		// if it's closer, user it
		if minPoint == -1 || v < minPoint {
			minPoint = v
			targetLocation = k
		}
	}

	return targetLocation
}
