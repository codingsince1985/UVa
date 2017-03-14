// UVa 10763 - Foreign Exchange

package main

import (
	"fmt"
	"os"
)

type location struct{ loc1, loc2 int }

func solve(locations []location) bool {
	locationMap := make(map[int]int)
	for _, loc := range locations {
		locationMap[loc.loc1]--
		locationMap[loc.loc2]++
	}
	for _, v := range locationMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("10763.in")
	defer in.Close()
	out, _ := os.Create("10763.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		locations := make([]location, n)
		for i := range locations {
			fmt.Fscanf(in, "%d%d", &locations[i].loc1, &locations[i].loc2)
		}
		if solve(locations) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
