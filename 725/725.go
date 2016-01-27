// UVa 725 - Division

package main

import (
	"fmt"
	"os"
)

func unique(n int) (bool, map[int]bool) {
	m := make(map[int]bool)
	for i := 0; i < 5; i++ {
		r := n % 10
		res := m[r]
		if res {
			return false, nil
		}
		m[r] = true
		n = n / 10
	}
	return true, m
}

func notIn(n int, m map[int]bool) bool {
	for i := 0; i < 5; i++ {
		r := n % 10
		res := m[r]
		if res {
			return false
		}
		n = n / 10
	}
	return true
}

func solve(n int) [][2]int {
	var ans [][2]int
	for i := 10234; i <= 98765; i++ {
		if ok, m := unique(i); ok {
			d := i / n
			if d * n != i {
				continue
			}
			if ok, _ := unique(d); ok {
				if notIn(d, m) {
					ans = append(ans, [2]int{i, d})
				}
			}
		}
	}
	return ans
}

func main() {
	in, _ := os.Open("725.in")
	defer in.Close()
	out, _ := os.Create("725.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		ans := solve(n)
		if len(ans) == 0 {
			fmt.Fprintf(out, "There are no solutions for %d.\n", n)
		} else {
			for i := range ans {
				fmt.Fprintf(out, "%d / %05d = %d\n", ans[i][0], ans[i][1], n)
			}
		}
		fmt.Fprintln(out)
	}
}
