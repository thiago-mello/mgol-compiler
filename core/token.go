package core

import "fmt"

type Token struct {
	Class  string
	Lexeme string
	Type   string
}

func (t *Token) String() string {
	return fmt.Sprintf("Classe: %s, Lexema: %s, Tipo: %s", t.Class, t.Lexeme, t.Type)
}
