package file

import (
	"testing"
)

// TestReadFileLinesFromNonExistingFile
// Given a non exisitng file
// When  ReadFileLines
// Then  should return an error
func TestReadFileLinesFromNonExistingFile(t *testing.T) {
	_, err := ReadFileLines("../../test/nonexisting.txt")
	if err == nil {
		t.Error("Didn't return an error when tried open non existing file")
	}
}

// TestReadFileLinesFromAnExistingFileShouldNotReturnError
// Given an existing file
// When  ReadFileLines
// Then  should not return an error
func TestReadFileLinesFromAnExistingFileShouldNotReturnError(t *testing.T) {
	_, err := ReadFileLines("../../test/emptyfile.txt")
	if err != nil {
		t.Errorf("Failed to open an existing file with error [%s]", err)
	}
}

// TestReadFileLinesFromEmptyFileShouldReturnAnEmptyFileStructure
// Given an empty file
// When ReadFileLines
// Then should return an empty File struct
func TestReadFileLinesFromEmptyFileShouldReturnAnEmptyFileStructure(t *testing.T) {
	f, _ := ReadFileLines("../../test/emptyfile.txt")
	if f.LastLineNumber != 0 && len(f.Lines) != 0 {
		t.Errorf(
			"Reading an empty file didn't return an empty structure, last line number [%d], number of lines [%d]",
			f.LastLineNumber,
			len(f.Lines),
		)
	}
}

// TestReadFileLinesFromAFileShouldReturnTheCorrectNumberOfLines
// Given an non empty file
// When  ReadFileLines
// Then  should return the correct number of lines
func TestReadFileLinesFromAFileShouldReturnTheCorrectNumberOfLines(t *testing.T) {
	f, _ := ReadFileLines("../../test/lusiadas.txt")
	if f.LastLineNumber != 10229 && len(f.Lines) != 10229 {
		t.Errorf(
			"Reading a file didn't return the correct structure, expected [10229], last line number [%d], number of lines [%d]",
			f.LastLineNumber,
			len(f.Lines),
		)
	}
}

// TestReadLineShouldReturnAnErrorWhenTheProvidedLineIsBiggerThanTheLastOne
// Given an line number bigger than the number of existing lines on a files
// When  ReadLine
// Then  should return an error
func TestReadLineShouldReturnAnErrorWhenTheProvidedLineIsBiggerThanTheLastOne(t *testing.T) {
	f, _ := ReadFileLines("../../test/lusiadas.txt")
	_, err := f.ReadLine(128031)
	if err == nil {
		t.Errorf("Didn't return an error when the user provides a line number bigger than the last one")
	}
}

// TestReadLineShouldReturnTheRightLine
// Given an zero or a negative number as line number
// When  ReadLine
// Then  should return an error
func TestReadLineShouldReturnErrorForNegativeOrZeroLineNumber(t *testing.T) {
	f, _ := ReadFileLines("../../test/lusiadas.txt")
	_, err := f.ReadLine(0)
	if err == nil {
		t.Error("Should return error for zero or negative line number")
	}
}

// TestReadLineShouldReturnTheRightLine
// Given an existing line number
// When  ReadLine
// Then  should return the right string
func TestReadLineShouldReturnTheRightLine(t *testing.T) {
	f, _ := ReadFileLines("../../test/lusiadas.txt")
	line, _ := f.ReadLine(10229)
	expected := "subscribe to our email newsletter to hear about new eBooks."
	if line != expected {
		t.Errorf("Failed to return the right line\n Expected [%s], result [%s]", expected, line)
	}
}
