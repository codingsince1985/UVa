// UVa 10420 - List of Conquests

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in, _ := os.Open("10420.in")
	defer in.Close()
	out, _ := os.Create("10420.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n int
	dict := make(map[string]int)
	for fmt.Fscanf(in, "%d", &n); n > 0 && s.Scan(); n-- {
		line := s.Text()
		dict[strings.Fields(line)[0]]++
	}
	var countries []string
	for k := range dict {
		countries = append(countries, k)
	}
	sort.Strings(countries)
	for _, v := range countries {
		fmt.Fprintln(out, v, dict[v])
	}
}
