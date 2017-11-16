package Day201611

import "strings"

type component struct {
	floor       int
	element     string
	isMicrochip bool
}

type componentList []component

func (slice componentList) Len() int {
	return len(slice)
}

func (slice componentList) Less(i, j int) bool {
	if slice[i].element == slice[j].element {
		return slice[i].isMicrochip
	}
	return slice[i].element < slice[j].element
}

func (slice componentList) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (c component) equals(cmp component) bool {
	return c.floor == cmp.floor &&
		c.element == cmp.element &&
		c.isMicrochip == cmp.isMicrochip
}

func (c component) getInitial() string {
	return strings.ToUpper(c.element[:1])
}

func (c component) shortName() string {
	if c.isMicrochip {
		return c.getInitial() + "M"
	}

	return c.getInitial() + "G"
}
