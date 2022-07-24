package actions

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/thiago-mello/mgol-compiler/core"
	"github.com/thiago-mello/mgol-compiler/parser/grammar"
	"github.com/thiago-mello/mgol-compiler/parser/parse_table"
)

type ParseAction struct {
	Action string
	State  int
	Rule   grammar.GrammarRule
}

var actionsMatrix = getActionsMatrix()

func getActionsMatrix() [][]string {
	filepath := "parser/parse_table/csv/actions.csv"

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("error opening actions.csv file")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		log.Fatal("could not read header from actions.csv")
	}

	actions, err := reader.ReadAll()
	if err != nil {
		log.Fatal("could not parse actions definition")
	}

	return actions
}

func parseActionString(action string) ParseAction {
	if action == "acc" {
		return ParseAction{
			Action: "accept",
		}
	}

	splitString := strings.Split(action, ".")
	actionNumber, err := strconv.Atoi(splitString[1])
	if err != nil {
		log.Fatal("invalid actions.csv file")
	}

	switch splitString[0] {
	case "s":
		return ParseAction{
			Action: "shift",
			State:  actionNumber,
		}

	case "r":
		return ParseAction{
			Action: "reduce",
			Rule:   grammar.GetRule(actionNumber),
		}

	default:
		return ParseAction{}
	}
}

func Action(state int, token core.Token) (ParseAction, bool) {
	var isInvalidState = state < 0 || state >= len(actionsMatrix)
	if isInvalidState {
		return ParseAction{}, false
	}

	columIndex, ok := parse_table.GetTokenIndex(token)
	if !ok {
		return ParseAction{}, false
	}
	action := actionsMatrix[state][columIndex]

	return parseActionString(action), true
}
