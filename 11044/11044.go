// UVa 11044 - Searching for Nessy

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11044.in")
	defer in.Close()
	out, _ := os.Create("11044.out")
	defer out.Close()

	var t, n, m int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d%d", &n, &m)
		fmt.Fprintln(out, int(n/3)*int(m/3))
	}
}
