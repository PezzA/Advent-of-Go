package Day201611

type component struct {
	Element string
	IsChip  bool
}

type floor []component

type facility struct {
	Lift   int
	Floors []floor
}

type stateList []facility
