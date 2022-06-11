package files

import (
	"bufio"
	"os"
)

func SourceReader(filePath string) (*bufio.Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	return reader, nil
}
