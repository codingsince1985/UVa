// UVa 11661 - Burger Time?

package main

import (
	"fmt"
	"os"
	"strings"
)

const max = 2000000

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getDistance(highway string) int {
	distance, drugIdx, restIdx := max, -1, -1
	for i := range highway {
		switch highway[i] {
		case 'D':
			drugIdx = i
			if restIdx != -1 {
				distance = min(distance, i-restIdx)
			}
		case 'R':
			restIdx = i
			if drugIdx != -1 {
				distance = min(distance, i-drugIdx)
			}
		}
	}
	return distance
}

func solve(highway string) int {
	if strings.Contains(highway, "Z") {
		return 0
	}
	return getDistance(highway)
}

func main() {
	in, _ := os.Open("11661.in")
	defer in.Close()
	out, _ := os.Create("11661.out")
	defer out.Close()

	var l int
	var highway string
	for {
		if fmt.Fscanf(in, "%d", &l); l == 0 {
			break
		}
		fmt.Fscanf(in, "%s", &highway)
		fmt.Fprintln(out, solve(highway))
	}
}
