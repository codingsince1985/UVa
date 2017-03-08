// UVa 106 - Fermat vs. Pythagoras

package main

import (
	"fmt"
	"os"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	in, _ := os.Open("106.in")
	defer in.Close()
	out, _ := os.Create("106.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		cache := make(map[int]bool)
		cnt := 0
		for i := 1; i*i <= n; i++ {
			for j := i + 1; j*j <= n; j += 2 {
				if gcd(i, j) == 1 {
					a := j*j - i*i
					b := 2 * i * j
					c := i*i + j*j
					if c > n {
						break
					}
					cnt++
					for times := 1; c*times <= n; times++ {
						cache[a*times], cache[b*times], cache[c*times] = true, true, true
					}
				}
			}
		}
		fmt.Fprintln(out, cnt, n-len(cache))
	}
}
