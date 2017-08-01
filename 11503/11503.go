// UVa 11503 - Virtual Friends

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

func solve(out *os.File, n int, friends [][2]int) {
	f, v := make([]int, n), make([]int, n)
	for i := range f {
		f[i], v[i] = i, 1
	}
	for _, fv := range friends {
		f1, f2 := unionFind(fv[0], f), unionFind(fv[1], f)
		if f1 != f2 {
			f[f1] = f2
			v[f2] += v[f1]
		}
		fmt.Fprintln(out, v[f2])
	}
}

func main() {
	in, _ := os.Open("11503.in")
	defer in.Close()
	out, _ := os.Create("11503.out")
	defer out.Close()

	var kase, n int
	var f1, f2 string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		friendMap := make(map[string]int)
		var friends [][2]int
		var count int
		for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
			fmt.Fscanf(in, "%s%s", &f1, &f2)
			if _, ok := friendMap[f1]; !ok {
				friendMap[f1] = count
				count++
			}
			if _, ok := friendMap[f2]; !ok {
				friendMap[f2] = count
				count++
			}
			friends = append(friends, [2]int{friendMap[f1], friendMap[f2]})
		}
		solve(out, count, friends)
	}
}
