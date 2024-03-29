package Day201922

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2019, 22, "Slam Shuffle", 6
}

func (td dayEntry) PuzzleInput() string {
	return `deal with increment 2
cut 3310
deal with increment 13
cut -9214
deal with increment 14
deal into new stack
deal with increment 26
deal into new stack
deal with increment 62
cut -1574
deal with increment 74
cut -7102
deal with increment 41
cut 7618
deal with increment 70
cut 7943
deal into new stack
deal with increment 52
cut -3134
deal with increment 21
deal into new stack
deal with increment 20
deal into new stack
deal with increment 61
cut -2810
deal with increment 60
cut 3355
deal with increment 13
cut 3562
deal with increment 55
cut 2600
deal with increment 47
deal into new stack
cut -7010
deal with increment 34
cut 1726
deal with increment 61
cut 2805
deal with increment 39
cut 1907
deal into new stack
cut 3915
deal with increment 14
cut -6590
deal into new stack
deal with increment 73
deal into new stack
deal with increment 31
cut 1000
deal with increment 3
cut 8355
deal with increment 2
cut -5283
deal with increment 50
cut -7150
deal with increment 71
cut 6728
deal with increment 58
cut -814
deal with increment 14
cut -8392
deal with increment 71
cut 7674
deal with increment 46
deal into new stack
deal with increment 55
cut 7026
deal with increment 17
cut 1178
deal with increment 10
cut -8205
deal with increment 27
cut -55
deal with increment 44
cut -2392
deal into new stack
cut 7385
deal with increment 36
cut -399
deal with increment 74
cut 6895
deal with increment 20
cut 4346
deal with increment 15
cut -4088
deal with increment 3
cut 1229
deal with increment 59
cut 4708
deal with increment 62
cut 2426
deal with increment 30
cut 7642
deal with increment 73
cut 9049
deal into new stack
cut -3866
deal with increment 68
deal into new stack
cut 1407`
}
