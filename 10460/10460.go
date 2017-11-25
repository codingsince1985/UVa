// UVa 10460 - Find the Permuted String

package main

import (
	"fmt"
	"os"
)

func insert(slice []byte, pos int, value byte) []byte {
	slice = slice[0 : len(slice)+1]
	copy(slice[pos+1:], slice[pos:])
	slice[pos] = value
	return slice
}

func solve(str string, n int) string {
	n--
	l := len(str)
	indices := make([]int, l)
	for i := l; i > 0; i-- {
		indices[i-1] = n % i
		n /= i
	}
	permuted := make([]byte, 0, l)
	for i, idx := range indices {
		permuted = insert(permuted, idx, str[i])
	}
	return string(permuted)
}

func main() {
	in, _ := os.Open("10460.in")
	defer in.Close()
	out, _ := os.Create("10460.out")
	defer out.Close()

	var kase, n int
	var str string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &str)
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, solve(str, n))
	}
}
