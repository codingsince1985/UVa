// UVa 10082 - WERTYU

package main

import (
	"bytes"
	"fmt"
	"os"
)

func buildMap() map[byte]byte {
	keys := "`1234567890-=QWERTYUIOP[]\\ASDFGHJKL;'ZXCVBNM,./"
	dict := make(map[byte]byte)
	for idx := range keys {
		if idx != 0 {
			dict[keys[idx]] = keys[idx-1]
		}
	}
	return dict
}

func main() {
	in, _ := os.Open("10082.in")
	defer in.Close()
	out, _ := os.Create("10082.out")
	defer out.Close()

	dict := buildMap()
	var c byte
	var buf bytes.Buffer
	for {
		if _, err := fmt.Fscanf(in, "%c", &c); err != nil {
			break
		}
		if c == '\n' {
			fmt.Fprintln(out, buf.String())
			buf.Reset()
		} else {
			if nc, ok := dict[c]; !ok {
				buf.WriteByte(c)
			} else {
				buf.WriteByte(nc)
			}
		}
	}
}
