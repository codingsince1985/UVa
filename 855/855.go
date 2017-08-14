// UVa 855 - Lunch in Grid City

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	in, _ := os.Open("855.in")
	defer in.Close()
	out, _ := os.Create("855.out")
	defer out.Close()

	var t, s, a, f int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d%d%d", &s, &a, &f)
		streets, avenues := make([]int, f), make([]int, f)
		for i := range streets {
			fmt.Fscanf(in, "%d%d", &streets[i], &avenues[i])
		}
		sort.Ints(streets)
		sort.Ints(avenues)
		fmt.Fprintf(out, "(Street: %d, Avenue: %d)\n", streets[(f-1)/2], avenues[(f-1)/2])
	}
}
