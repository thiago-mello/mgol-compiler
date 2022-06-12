package alphabet

var alphabet = map[string]int{
	"D":  0,
	"L":  1,
	",":  2,
	";":  3,
	":":  4,
	".":  5,
	"!":  6,
	"?":  7,
	"\\": 8,
	"*":  9,
	"+":  10,
	"-":  11,
	"/":  12,
	"(":  13,
	")":  14,
	"{":  15,
	"}":  16,
	"[":  17,
	"]":  18,
	"<":  19,
	">":  20,
	"=":  21,
	"\"": 22,
}

func BelongsToAlphabet(char rune) bool {
	_, ok := alphabet[string(char)]

	return ok
}

func GetRuneIndex(char rune) (int, bool) {
	index, ok := alphabet[string(char)]

	return index, ok
}
