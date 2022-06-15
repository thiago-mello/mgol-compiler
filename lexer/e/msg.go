package e

import (
	"fmt"
)

const (
	ERROR_INVALID_SYMBOL = "caractere inválido no código fonte"
	ERROR_INVALID_WORD   = "palavra não reconhecida no código fonte"
	END_OF_FILE_REACHED  = "fim do arquivo"
	INVALID_STATE_INDEX  = "o indice do estado atual não é valido"
)

func GetErrorMsg(errorMessage string, row int, column int) string {
	return fmt.Sprintf("ERRO LÉXICO - %s, linha %d, coluna %d\n", errorMessage, row, column)
}
