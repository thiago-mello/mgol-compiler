package e

import (
	"fmt"
)

const (
	ERROR_INVALID_SYMBOL   = "caractere inválido no código fonte"
	ERROR_INVALID_WORD     = "palavra não reconhecida no código fonte"
	END_OF_FILE_REACHED    = "fim do arquivo"
	INVALID_STATE_INDEX    = "o indice do estado atual não é valido"
	ERROR_OPENING_FILE     = "erro ao abrir o arquivo solicitado"
	FILE_PATH_NOT_PROVIDED = "caminho do arquivo fonte não informado"
)

func GetErrorMsg(errorMessage string, row int, column int) string {
	return fmt.Sprintf("ERRO LÉXICO - %s, linha %d, coluna %d", errorMessage, row, column)
}
