// UVa 729 - The Hamming Distance Problem

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	n, h int
	out  *os.File
)

func dfs(length, ones int, bits []string) {
	if length > n || ones > h {
		return
	}
	if length == n && ones == h {
		fmt.Fprintln(out, strings.Join(bits, ""))
		return
	}
	dfs(length+1, ones, append(bits, "0"))
	dfs(length+1, ones+1, append(bits, "1"))
}

func main() {
	in, _ := os.Open("729.in")
	defer in.Close()
	out, _ = os.Create("729.out")
	defer out.Close()

	var kase int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d%d", &n, &h)
		dfs(0, 0, nil)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
