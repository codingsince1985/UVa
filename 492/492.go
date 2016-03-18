// UVa 492 - Pig-Latin

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	in, _ := os.Open("492.in")
	defer in.Close()
	out, _ := os.Create("492.out")
	defer out.Close()

	var c byte
	var started bool
	var buf bytes.Buffer
	var suffix string
	for {
		if _, err := fmt.Fscanf(in, "%c", &c); err != nil {
			break
		}
		if !started {
			if c > 'A' && c < 'Z' || c > 'a' && c < 'z' {
				started = true
				u := strings.ToUpper(string(c))
				if !(u == "A" || u == "E" || u == "I" || u == "O" || u == "U") {
					suffix = string(c)
				} else {
					buf.WriteByte(c)
				}
				suffix += "ay"
			} else {
				buf.WriteByte(c)
			}
		} else {
			if !(c > 'A' && c < 'Z' || c > 'a' && c < 'z') {
				started = false
				buf.WriteString(suffix)
				suffix = ""
			}
			buf.WriteByte(c)
		}
	}
	fmt.Fprintln(out, buf.String())
}
