// UVa 120 - Stacks of Flapjacks

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func find(pancake string, pancakes []string) int {
	for i, v := range pancakes {
		if pancake == v {
			return i
		}
	}
	return -1
}

func flip(pos int, pancakes []string) {
	for i := 0; i <= pos/2; i++ {
		pancakes[i], pancakes[pos-i] = pancakes[pos-i], pancakes[i]
	}
}

func process(idx int, sorted, pancakes []string) []string {
	var steps []string
	switch {
	case sorted[idx] == pancakes[idx]:
	case pancakes[0] == sorted[idx]:
		steps = append(steps, strconv.Itoa(len(pancakes)-idx))
		flip(idx, pancakes)
	default:
		maxPos := find(sorted[idx], pancakes)
		flip(maxPos, pancakes)
		steps = append(steps, strconv.Itoa(len(pancakes)-maxPos))
		flip(idx, pancakes)
		steps = append(steps, strconv.Itoa(len(pancakes)-idx))
	}
	return steps
}

func solve(pancakes []string) []string {
	sorted := make([]string, len(pancakes))
	copy(sorted, pancakes)
	sort.Strings(sorted)
	var steps []string
	for i := len(sorted) - 1; i >= 0; i-- {
		steps = append(steps, process(i, sorted, pancakes)...)
	}
	return append(steps, "0")
}

func main() {
	in, _ := os.Open("120.in")
	defer in.Close()
	out, _ := os.Create("120.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		pancakes := strings.Split(line, " ")
		steps := solve(pancakes)
		fmt.Fprintln(out, line)
		fmt.Fprintln(out, strings.Join(steps, " "))
	}
}
