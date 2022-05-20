package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	filename := "prideAndPrejudice"
	data, err := os.ReadFile(fmt.Sprintf("%s.txt", filename))

	if err != nil {
		panic(err)
	}

	buffer := shift(data)
	err = os.WriteFile(fmt.Sprintf("%s.txt", shift([]byte(filename))), buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func shift(b []byte) *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})

	for _, b := range b {
		isUppercase := b >= 'A' && b <= 'Z'
		isLowercase := b >= 'a' && b <= 'z'

		next := b

		if isUppercase || isLowercase {
			var z byte = 'z'
			if isUppercase {
				z = 'Z'
			}

			next = b + 2

			if next > z {
				next = next - 26
			}
		}
		buffer.WriteByte(next)
	}
	return buffer
}
