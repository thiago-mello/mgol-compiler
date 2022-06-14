package dfa

import "github.com/thiago-mello/mgol-compiler/lexer/dfa/classes"

var finalStates = map[int]string{
	1:  classes.LOGICAL_OPERATION,
	2:  classes.IDENTIFICATION_NAME,
	3:  classes.LOGICAL_OPERATION,
	4:  classes.LOGICAL_OPERATION,
	5:  classes.LOGICAL_OPERATION,
	6:  classes.LOGICAL_OPERATION,
	7:  classes.ASSIGNMENT_OPERATOR,
	8:  classes.ARITHMETIC_OPERATORS,
	9:  classes.OPEN_PARENTHESES,
	10: classes.CLOSE_PARENTHESES,
	12: classes.LITERAL_CONSTANT,
	13: classes.END_OF_FILE,
	15: classes.COMMENT,
	16: classes.NUMERIC_CONSTANT_INT,
	19: classes.NUMERIC_CONSTANT_FLOAT,
	21: classes.SEMICOLON,
	22: classes.COMMA,
	24: classes.NUMERIC_CONSTANT_INT,
	26: classes.NUMERIC_CONSTANT_FLOAT,
}

func GetStateLabel(state int) (string, bool) {
	label, ok := finalStates[state]

	return label, ok
}
