package parse_table

import "github.com/thiago-mello/mgol-compiler/core"

var actionIndexes = map[string]int{
	"state":     0,
	"inicio":    1,
	"varinicio": 2,
	"varfim":    3,
	"PT_V":      4,
	"id":        5,
	"VIR":       6,
	"inteiro":   7,
	"real":      8,
	"literal":   9,
	"leia":      10,
	"escreva":   11,
	"Lit":       12,
	"Num":       13,
	"RCB":       14,
	"OPM":       15,
	"se":        16,
	"AB_P":      17,
	"FC_P":      18,
	"entao":     19,
	"OPR":       20,
	"fimse":     21,
	"repita":    22,
	"fimrepita": 23,
	"fim":       24,
	"EOF":       25,
}

func GetTokenIndex(token core.Token) (int, bool) {
	index, ok := actionIndexes[token.Class]

	return index, ok
}
