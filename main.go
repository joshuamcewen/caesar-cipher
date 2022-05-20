package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	filename := "prideAndPrejudice"

	start := time.Now()
	data, err := os.ReadFile(fmt.Sprintf("%s.txt", filename))

	if err != nil {
		panic(err)
	}

	buffer := shift(data)
	err = os.WriteFile(fmt.Sprintf("%s.txt", shift([]byte(filename))), buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("Execution: %s\n", elapsed)
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
