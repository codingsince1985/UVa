// UVa 216 - Getting in Line

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	c      [][2]int
	v      []bool
	res, r []int
	seg, s []float64
	min    float64
)

func backtracking(cur int, length float64) {
	if cur == len(c) {
		if length < min {
			min = length
			copy(seg, s)
			copy(res, r)
		}
		return
	}

	for i := range c {
		if !v[i] {
			v[i] = true
			r[cur] = i
			var seg float64
			if cur != 0 {
				n1, n2 := c[r[cur-1]], c[r[cur]]
				dx, dy := n1[0]-n2[0], n1[1]-n2[1]
				seg = math.Sqrt(float64(dx*dx+dy*dy)) + 16
				s[cur-1] = seg
			}
			backtracking(cur+1, length+seg)
			v[i] = false
		}
	}
}

func output(out *os.File, kase int) {
	fmt.Fprintln(out, "**********************************************************")
	fmt.Fprintf(out, "Network #%d\n", kase)
	for i := range seg {
		n1 := c[res[i]]
		n2 := c[res[i+1]]
		fmt.Fprintf(out, "Cable requirement to connect (%d,%d) to (%d,%d) is %.2f feet.\n", n1[0], n1[1], n2[0], n2[1], seg[i])
	}
	fmt.Fprintf(out, "Number of feet of cable required is %.2f.\n", min)
}

func main() {
	in, _ := os.Open("216.in")
	defer in.Close()
	out, _ := os.Create("216.out")
	defer out.Close()

	var n, x, y int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		c = make([][2]int, n)
		for i := range c {
			fmt.Fscanf(in, "%d%d", &x, &y)
			c[i] = [2]int{x, y}
		}

		v = make([]bool, n)
		r = make([]int, n)
		res = make([]int, n)
		s = make([]float64, n)
		seg = make([]float64, n-1)
		min = math.MaxFloat64
		backtracking(0, 0)
		output(out, kase)
	}
}
