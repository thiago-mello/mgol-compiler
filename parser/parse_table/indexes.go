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
	"P":         1,
	"V":         2,
	"LV":        3,
	"D":         4,
	"L":         5,
	"TIPO":      6,
	"A":         7,
	"ES":        8,
	"ARG":       9,
	"CMD":       10,
	"LD":        11,
	"OPRD":      12,
	"COND":      13,
	"CAB":       14,
	"EXP_R":     15,
	"CP":        16,
	"R":         17,
	"CABR":      18,
	"CPR":       19,
}

func GetTokenIndex(token core.Token) (int, bool) {
	index, ok := actionIndexes[token.Class]

	return index, ok
}

func GetNonTerminalIndex(nonTerminal string) (int, bool) {
	index, ok := actionIndexes[nonTerminal]

	return index, ok
}
