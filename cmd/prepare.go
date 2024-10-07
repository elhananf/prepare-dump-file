package cmd

import (
	"github.com/elhananf/prepare-dump-file/fileutils"
	"github.com/elhananf/prepare-dump-file/transformer"
)

func PrepareDump(inputPath string, outputPath string) error {
	inputFile, inputFileError := fileutils.OpenFileForReading(inputPath)
	if inputFileError != nil {
		return inputFileError
	}
	defer inputFile.Close()

	inputLines, fileReadError := fileutils.ReadLinesFromFile(inputFile)
	if fileReadError != nil {
		return fileReadError
	}

	outputLines := transformer.TransformLines(inputLines)

	outputFile, outputFileError := fileutils.OpenFileForWriting(outputPath)
	if outputFileError != nil {
		return outputFileError
	}
	defer outputFile.Close()

	fileWriteError := fileutils.WriteLinesToFile(outputLines, outputFile)
	if fileWriteError != nil {
		return fileWriteError
	}

	return nil
}
