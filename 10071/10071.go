// UVa 10071 - Back to High School Physics

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10071.in")
	defer in.Close()
	out, _ := os.Create("10071.out")
	defer out.Close()

	var v, t int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &v, &t); err != nil {
			break
		}
		fmt.Fprintln(out, 2*v*t)
	}
}
