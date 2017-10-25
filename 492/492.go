// UVa 492 - Pig-Latin

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const vowels = "AEIOU"

func isAlphabetic(b byte) bool { return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' }

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
			if isAlphabetic(c) {
				started = true
				if sc := string(c); !strings.Contains(vowels, strings.ToUpper(sc)) {
					suffix = sc
				} else {
					buf.WriteByte(c)
				}
				suffix += "ay"
			} else {
				buf.WriteByte(c)
			}
		} else {
			if !isAlphabetic(c) {
				started = false
				buf.WriteString(suffix)
				suffix = ""
			}
			buf.WriteByte(c)
		}
	}
	fmt.Fprintln(out, buf.String())
}
