package linereader

import (
	"bytes"
	"io"
)

type delimitedLineReader struct {
	reader     io.Reader
	delimiters [][]byte
	buffer     []byte
	done       bool
}

func (lr *delimitedLineReader) findNextDelimiter(data []byte) (int, []byte) {

	// If the data is empty, return nothing
	if len(data) == 0 {
		return -1, nil
	}

	// Loop through the delimeters
	for _, delim := range lr.delimiters {
		index := bytes.Index(data, delim)
		if index > -1 {
			return index, delim
		}
	}

	// Otherwise return nothing
	return -1, nil

}

func (lr *delimitedLineReader) scanLine() []byte {

	// Check if there is a delimiter
	foundIndex, foundDelim := lr.findNextDelimiter(lr.buffer)
	if foundIndex > -1 {

		// Get the found line
		foundLine := make([]byte, foundIndex)
		copy(foundLine, lr.buffer[:foundIndex])

		// Shift everything over by the appropriate number of bytes
		offset := len(foundLine) + len(foundDelim)
		copy(lr.buffer, lr.buffer[offset:])
		lr.buffer = lr.buffer[:len(lr.buffer)-offset]

		// Return the line that was found
		return foundLine

	}

	// Return nil otherwise
	return nil

}

func (lr *delimitedLineReader) Line() ([]byte, error) {

	// If the reader is done already, skip
	if lr.done {
		return nil, nil
	}

	// A buffer used for incremental reading of data from the source reader
	var buffer []byte

	for {

		// Find the first line in the accumulated buffer
		foundLine := lr.scanLine()
		if foundLine != nil {
			return foundLine, nil
		}

		// Initialize the buffer the first time around. We do this here conditionally
		// instead of once unconditionally to avoid allocating unused buffers
		if buffer == nil {
			buffer = make([]byte, 1024)
		}

		// Read some data into the buffer
		n, err := lr.reader.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}

		// If some bytes were read, copy it into the accumulated buffer
		if n > 0 {
			if len(lr.buffer) == 0 {
				lr.buffer = make([]byte, n)
				copy(lr.buffer, buffer[:n])
			} else {
				lr.buffer = append(lr.buffer, buffer[:n]...)
			}
		}

		// If we reached the end of the file, return all of the data in the buffer
		if err == io.EOF {
			lr.done = true
			if lr.buffer == nil {
				return nil, nil
			} else {
				finalLine := lr.buffer
				lr.buffer = nil
				return finalLine, nil
			}
		}

	}

}
