package fileutils

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFileForWriting(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func OpenFileForReading(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadLinesFromFile(inputFile *os.File) ([]string, error) {
	var lines []string
	buffer := make([]byte, 0, 64*1024)
	scanner := bufio.NewScanner(inputFile)
	scanner.Buffer(buffer, 1024*1024)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func WriteLinesToFile(lines []string, outputFile *os.File) error {
	w := bufio.NewWriter(outputFile)

	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	w.Flush()

	return nil
}
