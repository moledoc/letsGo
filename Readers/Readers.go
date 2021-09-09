package reader

import "io"

// Exercise: Reader
type MyReader struct{}

func (r MyReader) Read(buf []byte) (n int, err error) {
	buf[0] = 'A'
	return 1, nil
}

// My understanding of this solution:
// You add 'A' to the first buffer byte and return buffer length 1 and error nil.
// When using this reader, it will read inifinite times 'A', size the buffer does not overfill (we are adding the element to the first byte)
// and we don't return error io.EOF.

// Exercise: rot13Reader
type Rot13Reader struct {
	R io.Reader
}

func (rot *Rot13Reader) Read(buf []byte) (n int, err error) {

	// Read the reader from the type rot13Reader
	// This shoulw populate buf with the contents of rot.r
	// so if we have
	// 	```go
	//	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// 	r := rot13Reader{s}
	// ```
	// then after reading rot.r.Read(buf) we have populated the buffer with "Lbh penpxrq gur pbqr!"
	n, err = rot.R.Read(buf)

	// Now lets look et each element in the buffer
	for i := 0; i < len(buf); i++ {
		// if the element is between A and Z (A -> 65)
		// then rotate the letter 13 places among uppercase letters (rot13 algorithm)
		if buf[i] >= 'A' && buf[i] <= 'Z' {
			buf[i] = 65 + (buf[i]-65+13)%26
			// if the element is between a and z (a -> 97)
			// then rotate the letter 13 places among lowercase letters (rot13 algorithm)
		} else if buf[i] >= 'a' && buf[i] <= 'z' {
			buf[i] = 97 + (buf[i]-97+13)%26
		}
	}
	// return the rot.r n and err values.
	// the sample solution I checked did not explicitly return them,
	// but I (currently) like to be more verbose
	return n, err
}
