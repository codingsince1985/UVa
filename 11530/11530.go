// UVa 11530 - SMS Typing

package main

import (
	"bufio"
	"fmt"
	"os"
)

var keys = map[byte]int{
	' ': 1,
	'a': 1, 'b': 2, 'c': 3,
	'd': 1, 'e': 2, 'f': 3,
	'g': 1, 'h': 2, 'i': 3,
	'j': 1, 'k': 2, 'l': 3,
	'm': 1, 'n': 2, 'o': 3,
	'p': 1, 'q': 2, 'r': 3, 's': 4,
	't': 1, 'u': 2, 'v': 3,
	'w': 1, 'x': 2, 'y': 3, 'z': 4,
}

func main() {
	in, _ := os.Open("11530.in")
	defer in.Close()
	out, _ := os.Create("11530.out")
	defer out.Close()

	var t int
	fmt.Fscanf(in, "%d", &t)

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for i := 1; i <= t && s.Scan(); i++ {
		line := s.Text()
		cnt := 0
		for j := range line {
			cnt += keys[line[j]]
		}
		fmt.Fprintf(out, "Case #%d: %d\n", i, cnt)
	}
}
