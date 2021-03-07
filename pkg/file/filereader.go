package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type (
	// File a structure that stores the file path, an index for its lines and
	// the last line number
	File struct {
		Path           string // Path file path
		LastLineNumber int    // LastLineNumber last line number
		Lines          []Line // Lines a slice of lines
	}

	// Line a structure that stores the start and finish index of a line
	Line struct {
		Start  int64 // Start where the line starts
		Finish int64 // Finish where the line finishes
	}
)

// ReadFileLines Gatherer all file lines given a path string to a file
// returns the File struck and an error if any
func ReadFileLines(path string) (File, error) {
	file, err := os.Open(path)
	if err != nil {
		return File{}, fmt.Errorf("Failed to open file at path [%s], with the error [%+v]", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	var startIndex int64 = 0

	lines := make([]Line, 0)
	for scanner.Scan() {
		line := scanner.Text() //drops \n
		endIndex := startIndex + int64(len(line))
		lines = append(lines, Line{startIndex, endIndex})
		startIndex = endIndex + 2 // (\n) +  (endIndex + 1)
		lineNumber++
	}

	return File{
		Path:           path,
		Lines:          lines,
		LastLineNumber: lineNumber,
	}, nil
}

// ReadLine reads a line from a file
// returns line string and an error if any
func (file File) ReadLine(lineNumber int) (string, error) {
	if lineNumber <= 0 {
		return "", errors.New("line number must be a positive number")
	}

	if lineNumber > file.LastLineNumber {
		return "", errors.New("line number above the last line of the file")
	}

	fp, err := os.Open(file.Path)
	if err != nil {
		return "", fmt.Errorf("failed to open file at path [%s], with the error [%+v]", file.Path, err)
	}
	defer fp.Close()

	lineStartEnd := file.Lines[lineNumber-1]
	size := lineStartEnd.Finish - lineStartEnd.Start
	line := make([]byte, size)
	_, err = fp.Seek(lineStartEnd.Start, 0)
	if err != nil {
		return "", errors.New("failed to seek to the start of the line")
	}

	_, err = fp.Read(line)
	if err != nil {
		return "", errors.New("failed to read the line")
	}

	return string(line), nil
}
