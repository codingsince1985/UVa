// UVa 10719 - Quotient Polynomial

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(out *os.File, k int, a1 []int) {
	a2 := make([]int, len(a1)-1)
	a2[0] = a1[0]
	fmt.Fprintf(out, "q(x): %d", a2[0])
	for i := 1; i < len(a2); i++ {
		a2[i] = k*a2[i-1] + a1[i]
		fmt.Fprintf(out, " %d", a2[i])
	}
	fmt.Fprintf(out, "\nr = %d\n\n", a2[len(a2)-1]*k+a1[len(a1)-1])
}

func main() {
	in, _ := os.Open("10719.in")
	defer in.Close()
	out, _ := os.Create("10719.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var tmp int
	for s.Scan() {
		k, _ := strconv.Atoi(s.Text())
		s.Scan()
		var a []int
		for r := strings.NewReader(s.Text()); ; {
			if _, err := fmt.Fscanf(r, "%d", &tmp); err != nil {
				break
			}
			a = append(a, tmp)
		}
		solve(out, k, a)
	}
}
