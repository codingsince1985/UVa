// UVa 10583 - Ubiquitous Religions

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

func main() {
	in, _ := os.Open("10583.in")
	defer in.Close()
	out, _ := os.Create("10583.out")
	defer out.Close()

	var n, m, s1, s2, kase int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		f := make([]int, n)
		for i := range f {
			f[i] = i
		}
		for ; m > 0; m-- {
			fmt.Fscanf(in, "%d%d", &s1, &s2)
			f1, f2 := unionFind(s1-1, f), unionFind(s2-1, f)
			if f1 != f2 {
				f[f1] = f2
			}
		}
		count := 0
		for k := range f {
			if f[k] == k {
				count++
			}
		}
		kase++
		fmt.Fprintf(out, "Case %d: %d\n", kase, count)
	}
}
