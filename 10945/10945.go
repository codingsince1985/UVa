// UVa 10945 - Mother bear

package main

import (
	"bufio"
	"fmt"
	"os"
)

func isAlphabet(b byte) bool { return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' }

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' + b - 'A'
	}
	return b
}

func isPalindrome(line string) bool {
	var chars []byte
	for i := range line {
		if isAlphabet(line[i]) {
			chars = append(chars, toLower(line[i]))
		}
	}
	for i := 0; i < len(chars)/2; i++ {
		if chars[i] != chars[len(chars)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("10945.in")
	defer in.Close()
	out, _ := os.Create("10945.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		if line := s.Text(); line != "DONE" {
			if isPalindrome(line) {
				fmt.Fprintln(out, "You won’t be eaten!")
			} else {
				fmt.Fprintln(out, "Uh oh..")
			}
		}
	}
}
