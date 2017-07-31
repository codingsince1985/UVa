// UVa 272 - TEX Quotes

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("272.in")
	defer in.Close()
	out, _ := os.Create("272.out")
	defer out.Close()

	var c byte
	var buf bytes.Buffer
	var closed bool
	for {
		if _, err := fmt.Fscanf(in, "%c", &c); err != nil {
			break
		}
		if c == '"' {
			if closed {
				buf.WriteString("''")
			} else {
				buf.WriteString("``")
			}
			closed = !closed
		} else {
			buf.WriteByte(c)
		}
	}
	fmt.Fprintln(out, buf.String())
}
