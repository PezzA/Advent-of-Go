package common

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func MapToList(sourceMap map[rune]int) PairList {
	pl := make(PairList, len(sourceMap))

	i := 0
	for k, v := range sourceMap {
		pl[i] = Pair{Key: k, Value: v}
		i++
	}

	return pl
}

func (p PairList) Less(i, j int) bool {

	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}

	return p[i].Value < p[j].Value
}

func (p PairList) Len() int { return len(p) }

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
