package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (v MyReader) Read(s []byte) (int, error) {
	bytesWritten := 0
	for i := range s {
		s[i] = byte('A')
		bytesWritten++
	}
	return bytesWritten, nil
}


func main() {
	reader.Validate(MyReader{})
}
