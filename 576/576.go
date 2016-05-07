// UVa 576 - Haiku Review

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	vowels = map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'y': true,
	}
	syllables = []int{5, 7, 5}
)

func count(line string) int {
	contiguous := false
	cnt := 0
	for i := range line {
		if vowels[line[i]] {
			if !contiguous {
				cnt++
				contiguous = true
			}
		} else {
			contiguous = false
		}
	}
	return cnt
}

func main() {
	in, _ := os.Open("576.in")
	defer in.Close()
	out, _ := os.Create("576.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		str := s.Text()
		if str == "e/o/i" {
			break
		}
		lines := strings.Split(str, "/")
		done := false
		for i, v := range lines {
			if count(v) != syllables[i] {
				fmt.Fprintln(out, i+1)
				done = true
			}
		}
		if !done {
			fmt.Fprintln(out, "Y")
		}
	}
}
