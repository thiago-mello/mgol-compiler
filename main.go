package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/thiago-mello/mgol-compiler/files"
	"github.com/thiago-mello/mgol-compiler/lexer/e"
	"github.com/thiago-mello/mgol-compiler/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(e.FILE_PATH_NOT_PROVIDED)
	}

	filePath := os.Args[1]

	sourcereader, err := files.SourceReader(filePath)
	if err != nil {
		fmt.Println(errors.New(e.ERROR_OPENING_FILE))
		return
	}

	parser.Parse(sourcereader)
}
