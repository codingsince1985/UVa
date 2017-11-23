// UVa 347 - Run

package main

import (
	"fmt"
	"os"
	"strconv"
)

const max = 9876545

func initialize() []int {
	cache := make([]int, max)
	for i := max - 2; i >= 0; i-- {
		if runaround(i) {
			cache[i] = i
		} else {
			cache[i] = cache[i+1]
		}
	}
	return cache
}

func isValid(str string) bool {
	digitMap := map[byte]bool{'0': true}
	for i := range str {
		if digitMap[str[i]] {
			return false
		}
		digitMap[str[i]] = true
	}
	return true
}

func runaround(r int) bool {
	str := strconv.Itoa(r)
	if !isValid(str) {
		return false
	}
	l := len(str)
	for count, idx, visited := 0, 0, make(map[int]bool); count < l; count++ {
		curr := int(str[idx] - '0')
		if idx = (idx + curr) % l; visited[idx] {
			return false
		}
		visited[idx] = true
	}
	return true
}

func main() {
	in, _ := os.Open("347.in")
	defer in.Close()
	out, _ := os.Create("347.out")
	defer out.Close()

	cache := initialize()
	var r int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &r); r == 0 {
			break
		}
		fmt.Fprintf(out, "Case %d: %d\n", kase, cache[r])
	}
}
