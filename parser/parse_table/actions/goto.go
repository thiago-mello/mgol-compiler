package actions

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/thiago-mello/mgol-compiler/parser/parse_table"
)

func getGoToMatrix() [][]string {
	filepath := "parser/parse_table/csv/goto.csv"

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("error opening goto.csv file")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		log.Fatal("could not read header from goto.csv")
	}

	goTo, err := reader.ReadAll()
	if err != nil {
		log.Fatal("could not parse goto definition")
	}

	return goTo
}

var goToMatrix = getGoToMatrix()

func parseGoToString(gotoString string) int {
	if gotoString == "" {
		return -1
	}

	state, err := strconv.Atoi(gotoString)
	if err != nil {
		return -1
	}

	return state
}

func GoTo(state int, nonTerminal string) int {
	var isInvalidState = state < 0 || state >= len(goToMatrix)
	if isInvalidState {
		fmt.Println("Invalid goto state", state)
		return -1
	}

	columnIndex, ok := parse_table.GetNonTerminalIndex(nonTerminal)
	if !ok {
		fmt.Println("Invalid nonterminal '" + nonTerminal + "'")
		return -1
	}

	goTo := goToMatrix[state][columnIndex]

	return parseGoToString(goTo)
}
