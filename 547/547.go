// UVa 547 - DDF

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const max = 3000

func digitSum(n int) int {
	var sum int
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	return sum
}

func factorSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += digitSum(i)
		}
	}
	return sum
}

var cache = func() []int {
	cache := make([]int, max+1)
	for i := range cache {
		cache[i] = factorSum(i)
	}
	return cache
}()

func solve(m, n int) []string {
	var longest []string
	for i := m; i <= n; i++ {
		seq := []string{strconv.Itoa(i)}
		for curr, next := i, cache[i]; next != curr; next = cache[curr] {
			seq = append(seq, strconv.Itoa(next))
			curr = next
		}
		if len(seq) > len(longest) {
			longest = make([]string, len(seq))
			copy(longest, seq)
		}
	}
	return longest
}

func main() {
	in, _ := os.Open("547.in")
	defer in.Close()
	out, _ := os.Create("547.out")
	defer out.Close()

	var m, n int
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d%d", &m, &n); err != nil {
			break
		}
		fmt.Fprintf(out, "Input%d: %d %d\n", kase, m, n)
		fmt.Fprintf(out, "Output%d: %s\n", kase, strings.Join(solve(m, n), " "))
	}
}
