// UVa 11936 - The Lazy Lumberjacks

package main

import (
	"fmt"
	"os"
)

func triangle(a, b, c int) bool { return a > 0 && b > 0 && c > 0 && a+b > c && a+c > b && b+c > a }

func main() {
	in, _ := os.Open("11936.in")
	defer in.Close()
	out, _ := os.Create("11936.out")
	defer out.Close()

	var n, a, b, c int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d%d", &a, &b, &c)
		if triangle(a, b, c) {
			fmt.Fprintln(out, "OK")
		} else {
			fmt.Fprintln(out, "Wrong!!")
		}
	}
}
