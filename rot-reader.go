package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotReader rot13Reader) Read(b []byte) (int, error) {
	n, err := rotReader.r.Read(b)
	if len(b) == 0 {
		return 0, errors.New("buffer too small")
	}
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, err
}

func rot13(b byte) byte {
	if b > 64 && b < 123 {
		b += 13
		if b > 90 && b < 97 || b > 122 {
			b -= 26
		}
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
