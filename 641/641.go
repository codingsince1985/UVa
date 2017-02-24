// UVa 641 - Do the Untwist

package main

import (
	"fmt"
	"os"
)

func text2Code(c byte) int {
	switch c {
	case '_':
		return 0
	case '.':
		return 27
	default:
		return int(c - 'a' + 1)
	}
}

func code2Text(c int) byte {
	switch c {
	case 0:
		return '_'
	case 27:
		return '.'
	default:
		return byte('a' + c - 1)
	}
}

func solve(k int, line string) string {
	n := len(line)
	plain := make([]byte, n)
	for i := range line {
		plain[k*i%n] = code2Text((text2Code(line[i]) + i) % 28)
	}
	return string(plain)
}

func main() {
	in, _ := os.Open("641.in")
	defer in.Close()
	out, _ := os.Create("641.out")
	defer out.Close()

	var k int
	var line string
	for {
		if fmt.Fscanf(in, "%d", &k); k == 0 {
			break
		}
		fmt.Fscanf(in, "%s", &line)
		fmt.Fprintln(out, solve(k, line))
	}
}
