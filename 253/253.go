// UVa 253 - Cube painting

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(c1, c2 []string) bool {
	sort.Strings(c1[:2])
	sort.Strings(c2[:2])
	sort.Strings(c1[2:4])
	sort.Strings(c2[2:4])
	sort.Strings(c1[4:])
	sort.Strings(c2[4:])
	return strings.Join(c1, "") == strings.Join(c2, "")
}

func isSame(cube1, c2 []string) bool {
	c1 := make([]string, 6)
	copy(c1, cube1)
	if check(c1, c2) {
		return true
	}
	copy(c1, cube1)
	c1[0], c1[5], c1[1], c1[4] = c1[1], c1[4], c1[0], c1[5]
	if check(c1, c2) {
		return true
	}
	copy(c1, cube1)
	c1[0], c1[5], c1[2], c1[3] = c1[2], c1[3], c1[0], c1[5]
	if check(c1, c2) {
		return true
	}
	copy(c1, cube1)
	c1[1], c1[4], c1[2], c1[3] = c1[2], c1[3], c1[1], c1[4]
	return check(c1, c2)
}

func main() {
	in, _ := os.Open("253.in")
	defer in.Close()
	out, _ := os.Create("253.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		c1 := make([]string, 6)
		c2 := make([]string, 6)
		for i := range c1 {
			c1[i] = string(line[i])
			c2[i] = string(line[i+6])
		}
		if isSame(c1, c2) {
			fmt.Fprintln(out, "TRUE")
		} else {
			fmt.Fprintln(out, "FALSE")
		}
	}
}
