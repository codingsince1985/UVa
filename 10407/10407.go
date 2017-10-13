// UVa 10407 - Simple division

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func solve(lst []int) int {
	sort.Ints(lst)
	for i := 1; i < len(lst); i++ {
		lst[i] -= lst[0]
	}
	lst = lst[1:]
	remainder := lst[0]
	for i := 1; i < len(lst); i++ {
		remainder = gcd(remainder, lst[i])
	}
	return remainder
}

func main() {
	in, _ := os.Open("10407.in")
	defer in.Close()
	out, _ := os.Create("10407.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	var tmp int
	for s.Scan() {
		if line = s.Text(); line == "0" {
			break
		}
		var lst []int
		for r := strings.NewReader(line); ; {
			if fmt.Fscanf(r, "%d", &tmp); tmp == 0 {
				break
			}
			lst = append(lst, tmp)
		}
		fmt.Fprintln(out, solve(lst))
	}
}
