package lexer

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/thiago-mello/mgol-compiler/core"
	"github.com/thiago-mello/mgol-compiler/lexer/dfa"
	"github.com/thiago-mello/mgol-compiler/lexer/dfa/classes"
	"github.com/thiago-mello/mgol-compiler/lexer/e"
	"github.com/thiago-mello/mgol-compiler/symboltable"
)

func Scanner(sourceReader *bufio.Reader, row *int, column *int) (core.Token, bool) {
	var (
		currentState int
		lexeme       string
	)
	for {

		char, _, eofErr := sourceReader.ReadRune()
		if eofErr != nil && lexeme != "" { //EOF reached and lexeme not parsed yet
			token, ok := getToken(currentState, lexeme, *row, *column)

			sourceReader.UnreadRune()
			return token, ok
		}

		if eofErr != nil { //EOF
			return core.Token{Class: classes.END_OF_FILE, Lexeme: classes.END_OF_FILE}, true
		}

		tempState, stateTransitionError := dfa.StateTransition(currentState, char)

		if stateTransitionError != nil {
			defer fmt.Println(e.GetErrorMsg(stateTransitionError.Error(), *row, *column))

			return getToken(currentState, lexeme, *row, *column)
		}

		if tempState < 0 {
			if lexeme != "" {
				sourceReader.UnreadRune()
			}

			return getToken(currentState, lexeme, *row, *column)
		}

		currentState = tempState
		if currentState != 0 {
			lexeme += string(char)
		}
		(*column)++

		if isNewline(char) {
			(*row)++
			*column = 0
		}
	}
}

func getToken(currentState int, lexeme string, row int, column int) (core.Token, bool) {
	label, ok := dfa.GetStateLabel(currentState)
	if !ok {
		msg := e.GetErrorMsg(e.ERROR_INVALID_WORD, row, column)
		fmt.Println(msg)

		return core.Token{}, false
	}

	token, ok := parseState(label, lexeme)

	return token, ok
}

func parseState(label string, lexeme string) (core.Token, bool) {

	switch label {
	case classes.IDENTIFICATION_NAME:
		symbolToken, ok := symboltable.GetElement(lexeme)
		if !ok {
			returnToken := &core.Token{
				Class:  label,
				Lexeme: lexeme,
			}
			symboltable.AddOrUpdateIfExists(returnToken)

			return *returnToken, true
		}

		return *symbolToken, true
	case classes.COMMENT:
		return core.Token{}, false
	case classes.NUMERIC_CONSTANT_INT, classes.NUMERIC_CONSTANT_FLOAT:
		labelSplit := strings.Split(label, ".")

		return core.Token{
			Class:  labelSplit[0],
			Lexeme: lexeme,
			Type:   labelSplit[1],
		}, true
	case classes.LITERAL_CONSTANT:
		return core.Token{
			Class:  label,
			Lexeme: lexeme,
			Type:   "literal",
		}, true
	default:
		return core.Token{
			Class:  label,
			Lexeme: lexeme,
		}, true
	}
}

func isNewline(char rune) bool {
	return char == '\n'
}
