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
	"github.com/thiago-mello/mgol-compiler/utils/colors"
)

func Parse(reader *bufio.Reader) {
	row, column := new(int), new(int)
	*row = 1

	var stateStack stack.Stack
	const initialState = 0

	a := getNextValidToken(reader, row, column)
	stateStack.Push(initialState)

	accepted := false
	printedError := false

	for !accepted {
		s, ok := stateStack.PeekTop()
		if !ok {
			log.Fatal(e.ERROR_STACK_INDEX)
		}

		action, ok := actions.Action(s, a)
		if !ok {
			log.Fatal(e.ERROR_ACTIONS_INDEX)
		}

		switch action.Action {
		case "shift":
			stateStack.Push(action.State)
			a = getNextValidToken(reader, row, column)
			printedError = false
		case "reduce":
			stateStack.PopMultiple(action.Rule.NumberOfSymbols)
			t, ok := stateStack.PeekTop()
			if !ok {
				log.Fatal(e.ERROR_STACK_INDEX)
			}

			goTo := actions.GoTo(t, action.Rule.Left)
			stateStack.Push(goTo)

			fmt.Println(action.Rule.Rule)
			printedError = false

		case "accept":
			accepted = true

		case "error":
			stateStack.Push(action.State)
			if !printedError {
				colors.PrintColored(e.GetErrorMsg(s, *row, *column), colors.RED)
				printedError = true
			}

		default:
			if !printedError {
				colors.PrintColored(e.GetErrorMsg(s, *row, *column), colors.RED)
				printedError = true
			}
			a = getNextValidToken(reader, row, column)
			if a.Class == "EOF" {
				colors.PrintColored(e.ERROR_PANIC_END, colors.RED)
				accepted = true
			}
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
