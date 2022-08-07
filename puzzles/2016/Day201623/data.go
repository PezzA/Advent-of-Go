package Day201623

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string, int) {
	return 2016, 23, "Safe Cracking", 2
}

func (td dayEntry) PuzzleInput() string {
	return `cpy a b
dec b
cpy a d
cpy 0 a
cpy b c
inc a
dec c
jnz c -2
dec d
jnz d -5
dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 94 c
jnz 80 d
inc a
inc d
jnz d -2
inc c
jnz c -5`
}
