// UVa 598 - Bundling Newspapers

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	out        *os.File
	n, size    int
	newspapers []string
)

func dfs(level, now int, path []string) {
	if level == size {
		fmt.Fprintln(out, strings.Join(path, ", "))
		return
	}
	for i := now; i < n; i++ {
		dfs(level+1, i+1, append(path, newspapers[i]))
	}
}

func parse(line string) (int, int) {
	if line == "*" {
		return 1, n
	}
	tokens := strings.Split(line, " ")
	size1, _ := strconv.Atoi(tokens[0])
	if len(tokens) == 1 {
		return size1, size1
	}
	size2, _ := strconv.Atoi(tokens[1])
	return size1, size2
}

func main() {
	in, _ := os.Open("598.in")
	defer in.Close()
	out, _ = os.Create("598.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var newspaper string
	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	for s.Scan(); kase > 0 && s.Scan(); kase-- {
		line := s.Text()
		newspapers = nil
		for s.Scan() {
			if newspaper = s.Text(); newspaper == "" {
				break
			}
			newspapers = append(newspapers, newspaper)
		}
		n = len(newspapers)
		for s1, s2 := parse(line); s1 <= s2; s1++ {
			fmt.Fprintf(out, "Size %d\n", s1)
			size = s1
			dfs(0, 0, nil)
			fmt.Fprintln(out)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
