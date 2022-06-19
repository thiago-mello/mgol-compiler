package lexer

import (
	"testing"

	"github.com/thiago-mello/mgol-compiler/lexer/dfa/classes"
)

func TestInvalidFinalState(t *testing.T) {
	currentState := 11
	lexeme := "test"

	_, got := getToken(currentState, lexeme, 0, 0)
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestValidFinalState(t *testing.T) {
	currentState := 10
	lexeme := ")"

	token, ok := getToken(currentState, lexeme, 0, 0)
	want := true

	if ok != want {
		t.Errorf("got %v, want %v", ok, want)
	}

	if token.Class != classes.CLOSE_PARENTHESES {
		t.Errorf("invalid token returned")
	}

}

func TestCommentClassReturn(t *testing.T) {
	currentState := 15
	lexeme := "{comment}"

	_, got := getToken(currentState, lexeme, 0, 0)
	want := false

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNummericTestReturn(t *testing.T) {
	currentState := 16
	lexeme := "123"

	token, got := getToken(currentState, lexeme, 0, 0)
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if token.Class != "Num" || token.Type != "inteiro" {
		t.Error("Invalid token return")
	}
}

func TestLiteralConstant(t *testing.T) {
	currentState := 12
	lexeme := "\"constante literal\""

	token, got := getToken(currentState, lexeme, 0, 0)
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if token.Class != classes.LITERAL_CONSTANT {
		t.Error("Invalid token return")
	}
}

func TestKeywordReturn(t *testing.T) {
	currentState := 2
	lexeme := "inicio"

	token, got := getToken(currentState, lexeme, 0, 0)
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if token.Class != "inicio" || token.Type != "inicio" {
		t.Error("Invalid token return")
	}
}

func TestIdReturn(t *testing.T) {
	currentState := 2
	lexeme := "identificador1"

	token, got := getToken(currentState, lexeme, 0, 0)
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if token.Class != classes.IDENTIFICATION_NAME {
		t.Error("Invalid token return")
	}
}
