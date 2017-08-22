// UVa 165 - Stamps

package main

import (
	"fmt"
	"os"
)

var (
	values, tmp []int
	h, k, ans   int
)

func valid(curr, idx, num, total, last int) bool {
	switch {
	case total == curr:
		return true
	case idx == last || num == h:
		return false
	default:
		return valid(curr, idx+1, num, total, last) || valid(curr, idx, num+1, total+tmp[idx], last)
	}
}

func dfs(curr, last int) {
	if valid(curr, 0, 0, 0, last) {
		dfs(curr+1, last)
	}
	if last < k {
		tmp[last] = curr
		dfs(curr+1, last+1)
		tmp[last] = 0
	}
	if curr-1 > ans {
		ans = curr - 1
		copy(values, tmp)
	}
}

func main() {
	in, _ := os.Open("165.in")
	defer in.Close()
	out, _ := os.Create("165.out")
	defer out.Close()

	for {
		if fmt.Fscanf(in, "%d%d", &h, &k); h == 0 && k == 0 {
			break
		}
		values, tmp = make([]int, k), make([]int, k)
		ans = 0
		dfs(1, 0)
		for _, value := range values {
			fmt.Fprintf(out, "%3d", value)
		}
		fmt.Fprintf(out, " ->%3d\n", ans)
	}
}
