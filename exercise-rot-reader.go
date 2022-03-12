package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(out []byte) (int, error) {
	input := make([]byte, 8)
	inputWritten, inputError := r.r.Read(input)

	for i := 0; i < inputWritten; i++ {
		out[i] = rot13(input[i])
	}
	return inputWritten, inputError
}

func rot13(char byte) byte {
	rotatedChar := char + byte(13)
	if (isUppercase(char) && rotatedChar > byte('Z')) ||
		(!isUppercase(char) && rotatedChar > byte('z')){
		return rotatedChar - byte(26)
	}
	return rotatedChar
}

func isUppercase(char byte) bool {
	return char >= byte('A') && char <= byte('Z')
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
