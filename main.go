package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("prideAndPrejudice.txt")

	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer([]byte{})

	for _, b := range data {
		isUppercase := b > 'A' && b < 'Z'
		isLowercase := b > 'a' && b < 'z'

		next := b

		if isUppercase || isLowercase {
			var z byte = 'Z'
			if isLowercase {
				z = 'z'
			}

			next = b + 2

			if next > z {
				next = next - 26
			}
		}
		buffer.WriteByte(next)
	}

	fmt.Println(buffer.String())
}
