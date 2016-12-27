// UVa 524 - Prime Ring Problem

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	n       int
	visited []bool
	list    []int
	out     *os.File
)

func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if i*i > n {
			return true
		}
		if n%i == 0 {
			return false
		}
	}
	return true
}

func output() {
	str := make([]string, n)
	for i := range str {
		str[i] = strconv.Itoa(list[i+1])
	}
	fmt.Fprintln(out, strings.Join(str, " "))
}

func backtracking(curr int) {
	if curr == n+1 && isPrime(list[1]+list[n]) {
		output()
	}

	for i := 2; i <= n; i++ {
		if !visited[i] && isPrime(list[curr-1]+i) {
			visited[i] = true
			list[curr] = i
			backtracking(curr + 1)
			visited[i] = false
		}
	}
}

func main() {
	in, _ := os.Open("524.in")
	defer in.Close()
	out, _ = os.Create("524.out")
	defer out.Close()

	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}

		visited = make([]bool, n+1)
		list = make([]int, n+1)
		visited[1] = true
		list[1] = 1

		fmt.Fprintf(out, "Case %d:\n", kase)
		backtracking(2)
		fmt.Fprintln(out)
	}
}
