package parser

import (
	"bufio"
	"fmt"
	"log"

	"github.com/thiago-mello/mgol-compiler/core"
	"github.com/thiago-mello/mgol-compiler/lexer"
	"github.com/thiago-mello/mgol-compiler/parser/e"
	"github.com/thiago-mello/mgol-compiler/parser/parse_table/actions"
	"github.com/thiago-mello/mgol-compiler/parser/stack"
)

func Parse(reader *bufio.Reader) {
	row, column := new(int), new(int)

	var stateStack stack.Stack
	const initialState = 0

	a := getNextValidToken(reader, row, column)
	stateStack.Push(initialState)

	accepted := false

	for !accepted {
		s, ok := stateStack.PeekTop()
		if !ok {
			log.Fatal("cannot continue syntactic analisys due to invalid stack indexes")
		}

		action, ok := actions.Action(s, a)
		if !ok {
			log.Fatal("cannot continue syntactic analisys due to invalid action indexes")
		}

		switch action.Action {
		case "shift":
			stateStack.Push(action.State)
			a = getNextValidToken(reader, row, column)
		case "reduce":
			stateStack.PopMultiple(action.Rule.NumberOfSymbols)
			t, ok := stateStack.PeekTop()
			if !ok {
				log.Fatal("cannot continue syntactic analisys due to invalid stack indexes")
			}

			goTo := actions.GoTo(t, action.Rule.Left)
			stateStack.Push(goTo)

			fmt.Println(action.Rule.Rule)
		case "accept":
			accepted = true
		default:
			e.GetErrorMsg(s, *row, *column)
		}
	}

}

func getNextValidToken(reader *bufio.Reader, row *int, column *int) core.Token {
	for {
		token, ok := lexer.Scanner(reader, row, column)
		if ok {
			return token
		}
	}
}
