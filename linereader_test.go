package linereader_test

import (
	"strings"
	"testing"

	"github.com/connerdouglass/go-linereader"
)

func TestLineReader(t *testing.T) {
	type testCase struct {
		input  string
		delim  [][]byte
		output []string
	}
	testCases := []testCase{
		{"hello world!", Delims("e", "r"), []string{"h", "llo wo", "ld!"}},
	}
	for _, tc := range testCases {
		lr := linereader.WithDelimeters(
			strings.NewReader(tc.input),
			tc.delim,
		)
		output, err := linereader.ReadAllStrings(lr)
		if err != nil {
			t.Fatal(err.Error())
		}
		if !CompareResults(output, tc.output) {
			t.Fatalf(
				"input=\"%s\" expected=[%s], got [%s]",
				tc.input,
				strings.Join(tc.output, ","),
				strings.Join(output, ","),
			)
		}
	}
}

func Delims(strs ...string) [][]byte {
	byteSlices := make([][]byte, len(strs))
	for i := range byteSlices {
		byteSlices[i] = []byte(strs[i])
	}
	return byteSlices
}

func CompareResults(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
