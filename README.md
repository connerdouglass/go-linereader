# go-linereader

A small library for streaming lines from an `io.Reader`.

## Example

```go
// Create a default LineReader, delimited by newlines
lr := linereader.New( source /* io.Reader */ )
for {

    // Read a single line from the reader
    line, err := lr.Line()
    if err != nil {
        panic(err)
    }

    // A nil line means we've reached the end of the stream
    if line == nil {
        break
    }

    // Print the found line
    fmt.Println("LINE: ", string(line))

}
```

## Using custom delimiters

```go
lr := linereader.WithDelimiter(
    source, /* io.Reader */
    []byte("zz"),
)
```

The above example uses `zz` as a delimiter. Splitting `fizzbuzz123` by that delimiter would result in `[fi, bu, 123]`.

