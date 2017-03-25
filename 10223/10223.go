// UVa 10223 - How many nodes ?

package main

import (
	"fmt"
	"os"
)

var catalan = func() []uint64 {
	var i uint64
	c := []uint64{1}
	for i = 0; ; i++ {
		tmp := c[i] * (4*i + 2) / (i + 2)
		if tmp > 4294967295 {
			break
		}
		c = append(c, tmp)
	}
	c[0] = 0
	return c
}()

func main() {
	in, _ := os.Open("10223.in")
	defer in.Close()
	out, _ := os.Create("10223.out")
	defer out.Close()

	var n uint64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		for i, c := range catalan {
			if c == n {
				fmt.Fprintln(out, i)
			}
			if c >= n {
				break
			}
		}
	}
}
