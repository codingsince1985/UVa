// UVa 10499 - The Land of Justice

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10499.in")
	defer in.Close()
	out, _ := os.Create("10499.out")
	defer out.Close()

	var n int64
	for {
		if fmt.Fscanf(in, "%d", &n); n < 0 {
			break
		}
		if n == 0 {
			fmt.Fprintln(out, "0%")
		} else {
			fmt.Fprintf(out, "%d%%\n", 25*n)
		}
	}
}
