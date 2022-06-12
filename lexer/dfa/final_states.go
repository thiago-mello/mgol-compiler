package dfa

var finalStates = map[int]string{
	1:  "OPR",
	2:  "id",
	3:  "OPR",
	4:  "OPR",
	5:  "OPR",
	6:  "OPR",
	7:  "RCB",
	8:  "OPM",
	9:  "AB_P",
	10: "FC_P",
	12: "Lit",
	13: "EOF",
	15: "Comentario",
	16: "Num",
	19: "Num Real",
	21: "PT_V",
	22: "VIR",
	24: "Num",
	26: "Num Real",
}

func GetStateLabel(state int) (string, bool) {
	label, ok := finalStates[state]

	return label, ok
}
