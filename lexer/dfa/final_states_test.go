package dfa

import "testing"

func TestInvalidFinalState(t *testing.T) {
	_, ok := GetStateLabel(0)

	if ok {
		t.Error("invalid state returned ok")
	}
}
