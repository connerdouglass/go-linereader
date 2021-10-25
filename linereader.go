package linereader

import "io"

// LineReader is an interface that defines the Line() function, which reads a single line at a time
type LineReader interface {
	Line() ([]byte, error)
}

// New creates a new LineReader instance from a reader. The returned LineReader will read lines
// from the provided io.Reader using newlines as delimiters.
func New(reader io.Reader) LineReader {
	return WithDelimeter(
		reader,
		[]byte{'\n'},
	)
}

func WithDelimeter(reader io.Reader, delimeter []byte) LineReader {
	return WithDelimeters(
		reader,
		[][]byte{delimeter},
	)
}

func WithDelimeters(reader io.Reader, delimeters [][]byte) LineReader {
	return &delimitedLineReader{
		reader:     reader,
		delimiters: delimeters,
	}
}
