package main

import (
	"errors"
	"fmt"

	"github.com/thiago-mello/mgol-compiler/files"
	"github.com/thiago-mello/mgol-compiler/lexer"
	"github.com/thiago-mello/mgol-compiler/lexer/e"
)

func main() {
	row, column := new(int), new(int)
	var endOfFileReached bool = false

	*row = 1
	*column = 0

	sourcereader, err := files.SourceReader("./fonte.alg")
	if err != nil {
		fmt.Println(errors.New(e.ERROR_OPENING_FILE))
		return
	}

	for !endOfFileReached {
		token, ok, err := lexer.Scanner(sourcereader, row, column)
		if err != nil {
			if _, eofReached := err.(e.EndOfFileReachedError); eofReached {
				eofReached = true
			} else {
				fmt.Println(err.Error())
				continue
			}
		}

		if ok {
			fmt.Println(token.String())
		}
	}
}
