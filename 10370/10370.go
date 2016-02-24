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
	fmt.Fscanf(in, "%d", &c)
	for i := 0; i < c; i++ {
		fmt.Fscanf(in, "%d", &n)
		lst := make([]int, n)
		total := 0
		for j := range lst {
			fmt.Fscanf(in, "%d", &lst[j])
			total += lst[j]
		}
		avg := float64(total) / float64(n)
		cnt := 0
		for j := range lst {
			if float64(lst[j]) > avg {
				cnt++
			}
		}
		fmt.Fprintf(out, "%.3f%%\n", float64(cnt)*100/float64(n))
	}
}
