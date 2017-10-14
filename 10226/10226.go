// UVa 10226 - Hardwood Species

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func names(treeMap map[string]int) []string {
	name := make([]string, len(treeMap))
	var idx int
	for k := range treeMap {
		name[idx] = k
		idx++
	}
	sort.Strings(name)
	return name
}

func main() {
	in, _ := os.Open("10226.in")
	defer in.Close()
	out, _ := os.Create("10226.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var tree string
	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	for s.Scan(); kase > 0; kase-- {
		var count float64
		treeMap := make(map[string]int)
		for s.Scan() {
			if tree = s.Text(); tree == "" {
				break
			}
			treeMap[tree]++
			count++
		}
		for _, v := range names(treeMap) {
			fmt.Fprintf(out, "%s %.4f\n", v, float64(treeMap[v]*100.0)/count)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
