// UVa 10919 - Prerequisites?

package main

import (
	"fmt"
	"os"
)

type category struct {
	courseMap map[string]bool
	required  int
}

func solve(selected []string, categories []category) bool {
here:
	for i := range categories {
		var count int
		for _, course := range selected {
			if categories[i].courseMap[course] {
				if count++; count == categories[i].required {
					continue here
				}
			}
		}
		return false
	}
	return true
}

func main() {
	in, _ := os.Open("10919.in")
	defer in.Close()
	out, _ := os.Create("10919.out")
	defer out.Close()

	var k, m, c int
	var course string
	for {
		if fmt.Fscanf(in, "%d", &k); k == 0 {
			break
		}
		fmt.Fscanf(in, "%d", &m)
		selected := make([]string, k)
		for i := range selected {
			fmt.Fscanf(in, "%s", &selected[i])
		}
		categories := make([]category, m)
		for i := range categories {
			categories[i].courseMap = make(map[string]bool)
			for fmt.Fscanf(in, "%d%d", &c, &categories[i].required); c > 0; c-- {
				fmt.Fscanf(in, "%s", &course)
				categories[i].courseMap[course] = true
			}
		}
		if solve(selected, categories) {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}
