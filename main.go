package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/thiago-mello/mgol-compiler/files"
	"github.com/thiago-mello/mgol-compiler/lexer"
	"github.com/thiago-mello/mgol-compiler/lexer/dfa/classes"
	"github.com/thiago-mello/mgol-compiler/lexer/e"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(e.FILE_PATH_NOT_PROVIDED)
	}

	filePath := os.Args[1]
	row, column := new(int), new(int)
	var endOfFileReached bool = false

	*row = 1
	*column = 0

	sourcereader, err := files.SourceReader(filePath)
	if err != nil {
		fmt.Println(errors.New(e.ERROR_OPENING_FILE))
		return
	}

	for !endOfFileReached {
		token, ok := lexer.Scanner(sourcereader, row, column)

		if ok {
			fmt.Println(token.String())
		}

		if token.Class == classes.END_OF_FILE {
			endOfFileReached = true
		}
	}
}
