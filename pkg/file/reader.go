package file

type (
	// Reader an interface that represents a line reader
	Reader interface {
		// ReadLine reads a line from a file
		// returns line string and an error if any
		ReadLine(lineNumber int) (string, error)
	}
)
