// UVa 700 - Date Bugs

package main

import (
	"fmt"
	"os"
)

type computer struct{ y, a, b int }

func solve(computers []computer) int {
here:
	for n := 0; ; n++ {
		year := computers[0].y + n*(computers[0].b-computers[0].a)
		if year > 10000 {
			return -1
		}
		for _, computer := range computers[1:] {
			if (year-computer.y)%(computer.b-computer.a) != 0 {
				continue here
			}
		}
		return year
	}
}

func main() {
	in, _ := os.Open("700.in")
	defer in.Close()
	out, _ := os.Create("700.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		computers := make([]computer, n)
		for i := range computers {
			fmt.Fscanf(in, "%d%d%d", &computers[i].y, &computers[i].a, &computers[i].b)
		}
		fmt.Fprintf(out, "Case #%d:\n", kase)
		if year := solve(computers); year == -1 {
			fmt.Fprintln(out, "Unknown bugs detected.")
		} else {
			fmt.Fprintf(out, "The actual year is %d.\n", year)
		}
		fmt.Fprintln(out)
	}
}
