package symboltable

import "github.com/thiago-mello/mgol-compiler/core"

var reservedWords = [14]string{
	"inicio",
	"varinicio",
	"varfim",
	"escreva",
	"leia",
	"se",
	"entao",
	"fimse",
	"Repita",
	"fimRepita",
	"fim",
	"inteiro",
	"literal",
	"real",
}

var symbolTable = make(map[string]*core.Token, 100) //capacity of 100 to prevent early rehash

func init() {
	for _, word := range reservedWords {
		symbolTable[word] = &core.Token{Class: word, Lexeme: word, Type: word}
	}
}

func GetElement(lexeme string) (*core.Token, bool) {
	token, ok := symbolTable[lexeme]

	return token, ok
}

func AddOrUpdateIfExists(token *core.Token) {
	lexeme := token.Lexeme

	symbolTable[lexeme] = token
}
