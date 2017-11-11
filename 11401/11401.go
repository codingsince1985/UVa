// UVa 11401 - Triangle Counting

package main

import (
	"fmt"
	"os"
)

const max = 1000000

var f = func() []int64 {
	f := make([]int64, max+1)
	f[3] = 0
	for i := int64(4); i <= max; i++ {
		f[i] = f[i-1] + (i-2)*(i-2)/4
	}
	return f
}()

func main() {
	in, _ := os.Open("11401.in")
	defer in.Close()
	out, _ := os.Create("11401.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n < 3 {
			break
		}
		fmt.Fprintln(out, f[n])
	}
}
