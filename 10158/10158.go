// UVa 10158 - War

package main

import (
	"fmt"
	"os"
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func do(out *os.File, n, c, x, y int, f []int) {
	fx1, fx2 := unionFind(x, f), unionFind(x+n, f)
	fy1, fy2 := unionFind(y, f), unionFind(y+n, f)
	switch c {
	case 1:
		if fx1 == fy2 {
			fmt.Fprintln(out, -1)
		} else {
			f[fx1], f[fx2] = fy1, fy2
		}
	case 2:
		if fx1 == fy1 {
			fmt.Fprintln(out, -1)
		} else {
			f[fx1], f[fx2] = fy2, fy1
		}
	case 3:
		if fx1 == fy1 {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	case 4:
		if fx1 == fy2 {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}

func main() {
	in, _ := os.Open("10158.in")
	defer in.Close()
	out, _ := os.Create("10158.out")
	defer out.Close()

	var n, c, x, y int
	fmt.Fscanf(in, "%d", &n)
	f := make([]int, 2*n)
	for i := range f {
		f[i] = i
	}
	for {
		if fmt.Fscanf(in, "%d%d%d", &c, &x, &y); c == 0 && x == 0 && y == 0 {
			break
		}
		do(out, n, c, x, y, f)
	}
}
