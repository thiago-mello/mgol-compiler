package e

import "fmt"

var errors = map[int]string{
	0:  "Esperava a palavra 'inicio'",
	2:  "Esperava a palavra 'varinicio'",
	3:  "Esperava uma das palavras 'leia', 'escreva', 'se', 'repita', 'fim' ou um nome",
	4:  "Esperava uma das palavras 'inteiro', 'real', 'literal' ou 'varfim'",
	6:  "Esperava uma das palavras 'leia', 'escreva', 'se', 'repita', 'fim' ou um nome",
	7:  "Esperava uma das palavras 'leia', 'escreva', 'se', 'repita', 'fim' ou um nome",
	8:  "Esperava uma das palavras 'leia', 'escreva', 'se', 'repita', 'fim' ou um nome",
	9:  "Esperava uma das palavras 'leia', 'escreva', 'se', 'repita', 'fim' ou um nome",
	11: "Esperava um nome",
	12: "Esperava um nome, uma constante literal ou uma constante numérica",
	13: "Esperava o operador '<-'",
	14: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimse' ou um nome",
	15: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimrepita' ou um nome",
	16: "Esperava o símbolo '('",
	17: "Esperava o símbolo '('",
	19: "Esperava uma das palavras 'inteiro', 'real', 'literal' ou 'varfim'",
	20: "Esperava o símbolo ';'",
	21: "Esperava um nome",
	29: "Esperava o símbolo ';'",
	30: "Esperava o símbolo ';'",
	34: "Esperava um nome ou uma constante numérica",
	36: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimse' ou um nome",
	37: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimse' ou um nome",
	38: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimse' ou um nome",
	41: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimrepita' ou um nome",
	42: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimrepita' ou um nome",
	43: "Esperava uma das palavras 'leia', 'escreva', 'se', 'fimrepita' ou um nome",
	45: "Esperava um nome ou uma constante numérica",
	46: "Esperava um nome ou uma constante numérica",
	49: "Esperava o símbolo ';'",
	50: "Esperava o símbolo ','",
	53: "Esperava o símbolo ';'",
	54: "Esperava uma operação matemática",
	63: "Esperava o símbolo ')'",
	64: "Esperava uma operação lógica",
	65: "Esperava o símbolo ')'",
	67: "Esperava um nome",
	69: "Esperava um nome ou uma constante numérica",
	70: "Esperava a palavra 'entao'",
	71: "Esperava um nome ou uma constante numérica",
}

func GetErrorMsg(state int, row int, column int) string {
	expectedMessage := errors[state]

	if expectedMessage == "" {
		expectedMessage = "Símbolo inesperado"
	}

	return fmt.Sprintf("ERRO SINTÁTICO - %s na linha %d e coluna %d", expectedMessage, row, column)
}