package dfa

import (
	"testing"
)

type stateTransitionTest struct {
	state, expected int
	symbol          rune
}

var stateTests = []stateTransitionTest{
	{state: 0, symbol: 's', expected: 2},
	{state: 0, symbol: ',', expected: 22},
	{state: 12, symbol: '!', expected: -1},
	{state: 24, symbol: ' ', expected: -1},
	{state: 14, symbol: '3', expected: 14},
	{state: 16, symbol: '.', expected: 17},
}

var invalidStateTests = []stateTransitionTest{
	{state: 0, symbol: ':'},
	{state: 0, symbol: '['},
	{state: 12, symbol: '9'},
	{state: 24, symbol: ' '},
	{state: 15, symbol: '3'},
	{state: 16, symbol: 'v'},
}

func TestValidStateTransition(t *testing.T) {
	for _, test := range stateTests {
		value, _ := StateTransition(test.state, test.symbol)

		if value != test.expected {
			t.Errorf("state %d and symbol %c generated an invalid state transition", test.state, test.symbol)
		}
	}
}

func TestInvalidStateTransition(t *testing.T) {
	for _, test := range invalidStateTests {
		value, _ := StateTransition(test.state, test.symbol)

		if value != -1 {
			t.Error("invalid state transition did not return -1")
		}
	}
}

func TestTransitionErrors(t *testing.T) {
	_, err := StateTransition(0, '&')

	if err == nil {
		t.Error("Invalid symbol didn't raise an error")
	}

	_, err = StateTransition(len(transitionTable)+1, ' ')

	if err == nil {
		t.Error("Invalid state did not raise an error")
	}
}
