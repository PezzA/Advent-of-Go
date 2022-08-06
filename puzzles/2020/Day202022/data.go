package Day202022

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2020, 22, "Crab Combat"
}

func (td dayEntry) PuzzleInput() string {
	return `Player 1:
50
14
10
17
38
40
3
46
39
25
18
2
41
45
7
47
36
1
30
32
8
31
12
5
28

Player 2:
9
6
37
42
22
4
21
15
44
16
29
43
19
11
13
24
48
35
26
23
27
33
20
49
34`
}

func (td dayEntry) PuzzleSpec() string {
	return ``
}