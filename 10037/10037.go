// UVa 10037 - Bridge

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(p []int) (int, []string) {
	num := len(p)
	switch {
	case num > 3:
		t1 := p[0] + p[1]*2 + p[num-1]
		t2 := p[0]*2 + p[num-2] + p[num-1]
		time, steps := solve(p[:num-2])
		if t1 < t2 {
			curr := []string{fmt.Sprintf("%d %d\n%d\n%d %d\n%d", p[0], p[1], p[0], p[num-2], p[num-1], p[1])}
			return t1 + time, append(curr, steps...)
		}
		curr := []string{fmt.Sprintf("%d %d\n%d\n%d %d\n%d", p[0], p[num-2], p[0], p[0], p[num-1], p[0])}
		return t2 + time, append(curr, steps...)
	case num == 3:
		return p[0] + p[1] + p[2], []string{fmt.Sprintf("%d %d\n%d\n%d %d", p[0], p[1], p[0], p[0], p[2])}
	case num == 2:
		return p[1], []string{fmt.Sprintf("%d %d", p[0], p[1])}
	default:
		return p[0], []string{strconv.Itoa(p[0])}
	}
}

func main() {
	in, _ := os.Open("10037.in")
	defer in.Close()
	out, _ := os.Create("10037.out")
	defer out.Close()

	var kase, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d", &n)
		p := make([]int, n)
		for i := range p {
			fmt.Fscanf(in, "%d", &p[i])
		}
		sort.Ints(p)
		time, steps := solve(p)
		fmt.Fprintln(out, time)
		fmt.Fprintln(out, strings.Join(steps, "\n"))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
