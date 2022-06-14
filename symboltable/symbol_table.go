package symboltable

import "github.com/thiago-mello/mgol-compiler/lexer"

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

var symbolTable = make(map[string]*lexer.Token, 100) //capacity of 100 to prevent early rehash

func init() {
	for _, word := range reservedWords {
		symbolTable[word] = &lexer.Token{Class: word, Lexeme: word, Type: word}
	}
}

func GetElement(lexeme string) (*lexer.Token, bool) {
	token, ok := symbolTable[lexeme]

	return token, ok
}

func AddOrUpdateIfExists(token *lexer.Token) {
	lexeme := token.Lexeme

	symbolTable[lexeme] = token
}
