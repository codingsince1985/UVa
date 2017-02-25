// UVa 620 - Cellular Structure

package main

import (
	"fmt"
	"os"
	"strings"
)

func dfs(line string) bool {
	if len(line) == 0 {
		return false
	}
	switch {
	case len(line) == 1:
		return line == "A"
	case strings.HasPrefix(line, "B") && strings.HasSuffix(line, "A"):
		return dfs(line[1 : len(line)-1])
	case strings.HasSuffix(line, "AB"):
		return dfs(line[:len(line)-2])
	default:
		return false
	}
}

func solve(line string) string {
	if dfs(line) {
		switch {
		case len(line) == 1:
			return "SIMPLE"
		case strings.HasPrefix(line, "B") && strings.HasSuffix(line, "A"):
			return "MUTAGENIC"
		case strings.HasSuffix(line, "AB"):
			return "FULLY-GROWN"
		}
	}
	return "MUTANT"
}

func main() {
	in, _ := os.Open("620.in")
	defer in.Close()
	out, _ := os.Create("620.out")
	defer out.Close()

	var kase int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &line)
		fmt.Fprintln(out, solve(line))
	}
}
