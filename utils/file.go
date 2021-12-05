package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func ScanFile() (*bufio.Scanner, func(), error) {
	if len(os.Args) != 2 {
		return nil, nil, errors.New("missing filename input args parameter")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewScanner(file), func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}, nil
}
