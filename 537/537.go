// UVa 537 - Artificial Intelligence?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convert(f float64, unit byte) float64 {
	switch unit {
	case 'm':
		return f * .001
	case 'k':
		return f * 1000
	case 'M':
		return f * 1000000
	default:
		return f
	}
}

func parse(tokens [2]string) (float64, float64, float64) {
	var p, u, i, v float64
	var unit byte
	for _, token := range tokens {
		idx := strings.Index(token, "=")
		fmt.Sscanf(token[idx+1:], "%f%c", &v, &unit)
		v = convert(v, unit)
		switch token[idx-1] {
		case 'U':
			u = v
		case 'I':
			i = v
		case 'P':
			p = v
		}
	}
	return p, u, i
}

func solve(tokens [2]string) string {
	p, u, i := parse(tokens)
	switch {
	case p == 0:
		return fmt.Sprintf("P=%.2fW", u*i)
	case u == 0:
		return fmt.Sprintf("U=%.2fV", p/i)
	default:
		return fmt.Sprintf("I=%.2fA", p/u)
	}
}

func main() {
	in, _ := os.Open("537.in")
	defer in.Close()
	out, _ := os.Create("537.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var token string
	var n int
	fmt.Fscanf(in, "%d", &n)
	for i := 1; i <= n && s.Scan(); i++ {
		fmt.Fprintf(out, "Problem #%d\n", i)
		cnt := 0
		var tokens [2]string
		for r := strings.NewReader(s.Text()); ; {
			fmt.Fscanf(r, "%s", &token)
			if strings.Contains(token, "=") {
				tokens[cnt] = token
				if cnt++; cnt == 2 {
					break
				}
			}
		}
		fmt.Fprintf(out, "%s\n\n", solve(tokens))
	}
}
