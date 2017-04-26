// UVa 331 - Mapping the Swaps

package main

import (
	"fmt"
	"os"
	"sort"
)

var n, maps, steps int

func bubbleSort(num []int) int {
	steps := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
				steps++
			}
		}
	}
	return steps
}

func dfs(count int, num []int) {
	if count == steps {
		if sort.IntsAreSorted(num) {
			maps++
		}
		return
	}
	for i := 0; i < n-1; i++ {
		if num[i] > num[i+1] {
			num[i], num[i+1] = num[i+1], num[i]
			dfs(count+1, num)
			num[i], num[i+1] = num[i+1], num[i]
		}
	}
}

func solve(num []int) int {
	sorted := make([]int, n)
	copy(sorted, num)
	maps, steps = 0, bubbleSort(sorted)
	if !sort.IntsAreSorted(num) {
		dfs(0, num)
	}
	return maps
}

func main() {
	in, _ := os.Open("331.in")
	defer in.Close()
	out, _ := os.Create("331.out")
	defer out.Close()

	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		num := make([]int, n)
		for i := range num {
			fmt.Fscanf(in, "%d", &num[i])
		}
		fmt.Fprintf(out, "There are %d swap maps for input data set %d.\n", solve(num), kase)
	}
}
