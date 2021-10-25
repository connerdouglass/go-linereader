package linereader

func ReadAll(lr LineReader) ([][]byte, error) {

	// Fill a slice with all the lines
	var lines [][]byte
	for {

		// Read a line from the reader
		line, err := lr.Line()
		if err != nil {
			return nil, err
		}

		// If the line is nil, break from the loop
		if line == nil {
			break
		}

		// Add the line to the slice
		lines = append(lines, line)

	}

	// Return the lines
	return lines, nil

}

func ReadAllStrings(lr LineReader) ([]string, error) {

	// Read it into a slice of byte slices
	byteSlices, err := ReadAll(lr)
	if err != nil {
		return nil, err
	}

	// Convert each byte slice to a string
	lineStrings := make([]string, len(byteSlices))
	for i := range byteSlices {
		lineStrings[i] = string(byteSlices[i])
	}
	return lineStrings, nil

}
