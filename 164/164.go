// UVa 164 - String Computer

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

func trace(op [][]int) []string {
	var ed []string
	ed = append(ed, "E")
	l1, l2 := len(op)-1, len(op[0])-1
	for l1 != 0 && l2 != 0 {
		cur := op[l1][l2]
		switch cur {
		case 0:
			l1--
			l2--
			ed = append(ed, "M")
		case 1:
			l1--
			l2--
			ed = append(ed, "C")
		case 2:
			l1--
			ed = append(ed, "D")
		case 3:
			l2--
			ed = append(ed, "I")
		}
	}
	for l1 != 0 {
		l1--
		ed = append(ed, "D")
	}
	for l2 != 0 {
		l2--
		ed = append(ed, "I")
	}
	return ed
}

func make2dint(d1, d2 int) [][]int {
	s := make([][]int, d1+1)
	for i := range s {
		s[i] = make([]int, d2+1)
	}
	return s
}

func ed(s1, s2 string) []string {
	l1, l2 := len(s1), len(s2)
	dp := make2dint(l1+1, l2+1)
	op := make2dint(l1+1, l2+1)

	for i := 1; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= l2; j++ {
		dp[0][j] = j
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
				if k == 1 {
					op[i][j] = 1 // change
				} else {
					op[i][j] = 0 // move
				}
			case dp[i][j] == dp[i-1][j]+1:
				op[i][j] = 2 // delete
			default:
				op[i][j] = 3 // insert
			}
		}
	}
	return trace(op)
}

func edit(ops []string, s1, s2 string) []string {
	var ed []string
	p, p1, p2 := 1, 0, 0
	for i := len(ops) - 1; i >= 0; i-- {
		op := ops[i]
		var tmp string
		switch op {
		case "M":
			p++
			p1++
			p2++
		case "D":
			tmp = fmt.Sprintf("%s%c%02d", "D", s1[p1], p)
			p1++
		case "I":
			tmp = fmt.Sprintf("%s%c%02d", "I", s2[p2], p)
			p++
			p2++
		case "C":
			tmp = fmt.Sprintf("%s%c%02d", "C", s2[p2], p)
			p++
			p1++
			p2++
		case "E":
			tmp = "E"
		}
		ed = append(ed, tmp)
	}
	return ed
}

func output(out *os.File, ed []string) {
	for _, v := range ed {
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("164.in")
	defer in.Close()
	out, _ := os.Create("164.out")
	defer out.Close()

	var s1, s2 string
	for {
		if fmt.Fscanf(in, "%s", &s1); s1 == "#" {
			break
		}
		fmt.Fscanf(in, "%s", &s2)
		ops := ed(s1, s2)
		output(out, edit(ops, s1, s2))
	}
}
