// Package utils provides shared parsing and math helpers used across puzzle solutions.
package utils

import (
	"bufio"
	"io"
)

// ReadLines reads all lines from an io.Reader and returns them as a string slice.
func ReadLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
