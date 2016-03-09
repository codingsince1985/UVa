// UVa 401 - Palindromes

package main

import (
	"fmt"
	"os"
)

var dict map[byte]byte = func() map[byte]byte {
	srcBytes := []byte{'A', 'E', 'H', 'I', 'J', 'L', 'M', 'O', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '1', '2', '3', '5', '8'}
	dstBytes := []byte{'A', '3', 'H', 'I', 'L', 'J', 'M', 'O', '2', 'T', 'U', 'V', 'W', 'X', 'Y', '5', '1', 'S', 'E', 'Z', '8'}
	dict := make(map[byte]byte)
	for i := range srcBytes {
		dict[srcBytes[i]] = dstBytes[i]
	}
	return dict
}()

var isPalindromic = func(a, b byte) bool { return a == b }

var isMirrored = func(a, b byte) bool { return dict[a] == b }

func testIf(str string, testMode func(a, b byte) bool) bool {
	size := len(str)
	half := size / 2
	for i := 0; i < half; i++ {
		if !testMode(str[i], str[size-1-i]) {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("401.in")
	defer in.Close()
	out, _ := os.Create("401.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		p := testIf(line, isPalindromic)
		m := testIf(line, isMirrored)
		switch {
		case !p && !m:
			fmt.Fprintf(out, "%s -- is not a palindrome.\n\n", line)
		case p && !m:
			fmt.Fprintf(out, "%s -- is a regular palindrome.\n\n", line)
		case !p && m:
			fmt.Fprintf(out, "%s -- is a mirrored string.\n\n", line)
		default:
			fmt.Fprintf(out, "%s -- is a mirrored palindrome.\n\n", line)
		}
	}
}
