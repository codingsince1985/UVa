// UVa 325 - Identifying Legal Pascal Real Constants

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func checkLeft(s string, hasRight bool) bool {
	if idx := strings.Index(s, "."); idx >= 0 {
		return !(idx == len(s)-1 || !isInt(s[0:idx]) || !isInt(s[idx+1:]))
	}
	return hasRight
}

func checkRight(s string) bool { return !(strings.Contains(s, ".") || !isInt(s)) }

func check(line string) bool {
	line = strings.ToUpper(line)
	if strings.Contains(line, "E") {
		parts := strings.Split(line, "E")
		if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
			return false
		}
		return checkLeft(parts[0], true) && checkRight(parts[1])
	}
	return checkLeft(line, false)
}

func main() {
	in, _ := os.Open("325.in")
	defer in.Close()
	out, _ := os.Create("325.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "*" {
			break
		}
		line = strings.TrimSpace(line)
		if fmt.Fprintf(out, "%s ", line); check(line) {
			fmt.Fprintln(out, "is legal.")
		} else {
			fmt.Fprintln(out, "is illegal.")
		}
	}
}
