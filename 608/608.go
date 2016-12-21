// UVa 608 - Counterfeit Dollar

package main

import (
	"fmt"
	"os"
	"strings"
)

type weighing struct {
	left, right string
	result      int
}

var (
	diffs      = [2]int{-1, 1}
	balanceMap = map[string]int{"up": -1, "down": 1, "even": 0}
	weightMap  = map[int]string{-1: "light", 1: "heavy"}
)

func verify(w weighing, c string, diff int) bool {
	l, r := 0, 0
	if strings.Contains(w.left, c) {
		l = diff
	}
	if strings.Contains(w.right, c) {
		r = diff
	}
	return r-l == w.result
}

func solve(weighings [3]weighing) (string, int) {
	for c := 'A'; c <= 'L'; c++ {
	here:
		for _, diff := range diffs {
			for _, w := range weighings {
				if !verify(w, string(c), diff) {
					continue here
				}
			}
			return string(c), diff
		}
	}
	return "", 0
}

func main() {
	in, _ := os.Open("608.in")
	defer in.Close()
	out, _ := os.Create("608.out")
	defer out.Close()

	var kase int
	var left, right, result string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		var weighings [3]weighing
		for i := range weighings {
			fmt.Fscanf(in, "%s%s%s", &left, &right, &result)
			weighings[i] = weighing{left, right, balanceMap[result]}
		}
		c, diff := solve(weighings)
		fmt.Fprintf(out, "%s is the counterfeit coin and it is %s.\n", c, weightMap[diff])
	}
}
