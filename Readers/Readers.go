package reader

type MyReader struct{}

func (r MyReader) Read(buf []byte) (n int, err error) {
	buf[0] = 'A'
	return 1, nil
}

// My understanding of this solution:
// You add 'A' to the first buffer byte and return buffer length 1 and error nil.
// When using this reader, it will read inifinite times 'A', size the buffer does not overfill (we are adding the element to the first byte)
// and we don't return error io.EOF.
