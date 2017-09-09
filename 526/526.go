// UVa 526 - String Distance and Transform Process

package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func make2d(d1, d2 int) [][]int {
	s := make([][]int, d1)
	for i := range s {
		s[i] = make([]int, d2)
	}
	return s
}

func trace(op [][]int) []string {
	var e []string
	l1, l2 := len(op)-1, len(op[0])-1
	for l1 != 0 && l2 != 0 {
		cur := op[l1][l2]
		switch cur {
		case 0:
			l1--
			l2--
			e = append(e, "M")
		case 1:
			l1--
			l2--
			e = append(e, "C")
		case 2:
			l2--
			e = append(e, "I")
		case 3:
			l1--
			e = append(e, "D")
		}
	}

	for l1 != 0 {
		e = append(e, "D")
		l1--
	}
	for l2 != 0 {
		e = append(e, "I")
		l2--
	}
	return e
}

func edit(s1, s2 string) [][]int {
	l1, l2 := len(s1), len(s2)
	dp := make2d(l1+1, l2+1)
	op := make2d(l1+1, l2+1)

	for i := range dp {
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	var k int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] {
				k = 0
			} else {
				k = 1
			}
			dp[i][j] = min(dp[i-1][j-1]+k, min(dp[i-1][j]+1, dp[i][j-1]+1))
			switch {
			case dp[i][j] == dp[i-1][j-1]+k:
				if k == 0 {
					op[i][j] = 0 // move
				} else {
					op[i][j] = 1 // change
				}
			case dp[i][j] == dp[i][j-1]+1:
				op[i][j] = 2 // insert
			default:
				op[i][j] = 3 // delete
			}
		}
	}
	return op
}

func nonMove(e []string) int {
	count := 0
	for _, v := range e {
		if v != "M" {
			count++
		}
	}
	return count
}

func output(out *os.File, e []string, s2 string) {
	fmt.Fprintln(out, nonMove(e))
	count := 0
	p, p1, p2 := 1, 0, 0
	for i := len(e) - 1; i >= 0; i-- {
		switch e[i] {
		case "M":
			p++
			p1++
			p2++
		case "C":
			count++
			fmt.Fprintf(out, "%d Replace %d,%c\n", count, p, s2[p2])
			p++
			p1++
			p2++
		case "I":
			count++
			fmt.Fprintf(out, "%d Insert %d,%c\n", count, p, s2[p2])
			p++
			p2++
		case "D":
			count++
			fmt.Fprintf(out, "%d Delete %d\n", count, p)
			p1++
		}
	}
}

func main() {
	in, _ := os.Open("526.in")
	defer in.Close()
	out, _ := os.Create("526.out")
	defer out.Close()

	var s1, s2 string
	first := true
	for {
		if _, err := fmt.Fscanf(in, "%s", &s1); err != nil {
			break
		}
		fmt.Fscanf(in, "%s", &s2)
		e := trace(edit(s1, s2))
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		output(out, e, s2)
	}
}
