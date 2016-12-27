// UVa 10051 - Tower of Cubes

package main

import (
	"fmt"
	"os"
)

type sides struct{ front, back, left, right, top, bottom int }

var (
	side = []string{"front", "back", "left", "right", "top", "bottom"}
	out  *os.File
)

func getBottomColor(t int, c sides) int {
	switch t {
	case 0:
		return c.back
	case 1:
		return c.front
	case 2:
		return c.right
	case 3:
		return c.left
	case 4:
		return c.bottom
	default:
		return c.top
	}
}

func getUpperColor(t int, c sides) int {
	switch t {
	case 0:
		return c.front
	case 1:
		return c.back
	case 2:
		return c.left
	case 3:
		return c.right
	case 4:
		return c.top
	default:
		return c.bottom
	}
}

func solve(c []sides) ([][6][2]int, int, [2]int) {
	l := len(c)
	dp := make([][6]int, l)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	pre := make([][6][2]int, l)
	mx := 0
	var st [2]int

	for i := 1; i < l; i++ {
		for j := 0; j < 6; j++ {
			for k := 0; k < i; k++ {
				for m := 0; m < 6; m++ {
					if getUpperColor(j, c[i]) == getBottomColor(m, c[k]) {
						if dp[k][m]+1 > dp[i][j] {
							dp[i][j] = dp[k][m] + 1
							pre[i][j] = [2]int{k, m}
							if mx < dp[i][j] {
								mx = dp[i][j]
								st = [2]int{i, j}
							}
						}
					}
				}
			}
		}
	}
	return pre, mx, st
}

func output(pre [][6][2]int, mx int, st [2]int) {
	var res [][2]int
	res = append(res, st)
	for i := 0; i < mx-1; i++ {
		res = append(res, pre[st[0]][st[1]])
		st = pre[st[0]][st[1]]
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Fprintln(out, res[i][0]+1, side[res[i][1]])
	}
}

func main() {
	in, _ := os.Open("10051.in")
	defer in.Close()
	out, _ = os.Create("10051.out")
	defer out.Close()

	var n, s1, s2, s3, s4, s5, s6 int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		c := make([]sides, n)
		for i := range c {
			fmt.Fscanf(in, "%d%d%d%d%d%d", &s1, &s2, &s3, &s4, &s5, &s6)
			cube := sides{s1, s2, s3, s4, s5, s6}
			c[i] = cube
		}
		pre, mx, st := solve(c)
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Case #%d\n%d\n", kase, mx)
		output(pre, mx, st)
	}
}
