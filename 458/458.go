// UVa 458 - The Decoder

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("458.in")
	defer in.Close()
	out, _ := os.Create("458.out")
	defer out.Close()

	var c byte
	var buf bytes.Buffer
	var delta byte = '1' - '*'
	for {
		if _, err := fmt.Fscanf(in, "%c", &c); err != nil {
			break
		}
		if c == '\n' {
			fmt.Fprintln(out, buf.String())
			buf.Reset()
		} else {
			buf.WriteByte(c - delta)
		}
	}
}
