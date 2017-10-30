// UVa 551 - Nesting a Bunch of Brackets

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	o1, c1, o2, c2, o3, c3, o4, c4 = '(', ')', '[', ']', '{', '}', '<', '>'
	o, c                           = "(*", "*)"
)

var (
	openMap  = map[byte]bool{o1: true, o2: true, o3: true, o4: true}
	closeMap = map[byte]bool{c1: true, c2: true, c3: true, c4: true}
	matchMap = map[byte]string{c1: "(", c2: "[", c3: "{", c4: "<"}
)

func solve(line string) (int, bool) {
	var stack []string
	for count, idx := 1, 0; idx < len(line); count++ {
		switch {
		case strings.HasPrefix(line[idx:], o):
			stack = append([]string{o}, stack...)
			idx++
		case openMap[line[idx]]:
			stack = append([]string{string(line[idx])}, stack...)
		case strings.HasPrefix(line[idx:], c):
			if len(stack) == 0 || stack[0] != o {
				return count, false
			}
			stack = stack[1:]
			idx++
		case closeMap[line[idx]]:
			if len(stack) == 0 || stack[0] != matchMap[line[idx]] {
				return count, false
			}
			stack = stack[1:]
		}
		idx++
	}
	return -1, true
}

func main() {
	in, _ := os.Open("551.in")
	defer in.Close()
	out, _ := os.Create("551.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		if count, ok := solve(s.Text()); ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintf(out, "NO %d\n", count)
		}
	}
}
