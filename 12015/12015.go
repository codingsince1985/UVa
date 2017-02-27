// UVa 12015 - Google is Feeling Lucky

package main

import (
	"fmt"
	"os"
	"strings"
)

type website struct {
	url       string
	relevance int
}

func solve(websites []website, max int) []string {
	var highest []string
	for _, w := range websites {
		if w.relevance == max {
			highest = append(highest, w.url)
		}
	}
	return highest
}

func main() {
	in, _ := os.Open("12015.in")
	defer in.Close()
	out, _ := os.Create("12015.out")
	defer out.Close()

	var t int
	fmt.Fscanf(in, "%d", &t)
	for kase := 1; kase <= t; kase++ {
		var max int
		websites := make([]website, 10)
		for i := range websites {
			fmt.Fscanf(in, "%s%d", &websites[i].url, &websites[i].relevance)
			if websites[i].relevance > max {
				max = websites[i].relevance
			}
		}
		fmt.Fprintf(out, "Case #%d:\n", kase)
		fmt.Fprintln(out, strings.Join(solve(websites, max), "\n"))
	}
}
