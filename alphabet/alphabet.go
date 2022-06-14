package alphabet

import "unicode"

var alphabet = map[rune]int{
	'D':  0,
	'L':  1,
	'E':  2,
	'e':  3,
	',':  4,
	';':  5,
	':':  6,
	'.':  7,
	'!':  8,
	'?':  9,
	'\\': 10,
	'*':  11,
	'+':  12,
	'-':  13,
	'/':  14,
	'(':  15,
	')':  16,
	'{':  17,
	'}':  18,
	'[':  19,
	']':  20,
	'<':  21,
	'>':  22,
	'=':  23,
	'"':  24,
	'_':  25,
}

func GetRuneIndex(char rune) (int, bool) {
	char = checkForLetterAndDigit(char)
	index, ok := alphabet[char]

	return index, ok
}

func checkForLetterAndDigit(char rune) rune {
	switch {
	case char == 'E' || char == 'e':
		return char
	case unicode.IsLetter(char):
		return 'L'
	case unicode.IsDigit(char):
		return 'D'
	default:
		return char
	}
}
