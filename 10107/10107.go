// UVa 10107 - What is the Median?

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	in, _ := os.Open("10107.in")
	defer in.Close()
	out, _ := os.Create("10107.out")
	defer out.Close()

	var n int
	var lst []int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		lst = append(lst, n)
		sort.Ints(lst)
		size := len(lst)
		if size%2 == 0 {
			fmt.Fprintln(out, (lst[size/2-1]+lst[size/2])/2)
		} else {
			fmt.Fprintln(out, lst[size/2])
		}
	}
}
