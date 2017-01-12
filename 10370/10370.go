// UVa 10370 - Above Average

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10370.in")
	defer in.Close()
	out, _ := os.Create("10370.out")
	defer out.Close()

	var c, n int
	for fmt.Fscanf(in, "%d", &c); c > 0; c-- {
		fmt.Fscanf(in, "%d", &n)
		lst := make([]int, n)
		total := 0
		for i := range lst {
			fmt.Fscanf(in, "%d", &lst[i])
			total += lst[i]
		}
		avg := float64(total) / float64(n)
		cnt := 0
		for _, v := range lst {
			if float64(v) > avg {
				cnt++
			}
		}
		fmt.Fprintf(out, "%.3f%%\n", float64(cnt)*100/float64(n))
	}
}
