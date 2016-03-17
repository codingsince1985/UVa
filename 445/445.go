// UVa 445 - Marvelous Mazes

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextLine(s *bufio.Scanner) (string, bool) {
	if ok := s.Scan(); ok {
		return s.Text(), true
	}
	return "", false
}

func outputLine(out *os.File, line string) {
	cnt := 0
	for _, v := range []byte(line) {
		if v == 'b' {
			v = ' '
		}
		switch {
		case v >= '0' && v <= '9':
			tmp, _ := strconv.Atoi(string(v))
			cnt += tmp
		default:
			for i := 0; i < cnt; i++ {
				fmt.Fprintf(out, "%c", v)
			}
			cnt = 0
		}
	}
	fmt.Fprintln(out)
}

func solve(out *os.File, maze string) {
	if maze[len(maze)-1] == '!' {
		maze = maze[:len(maze)-1]
	}
	for _, v := range strings.Split(maze, "!") {
		outputLine(out, v)
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("445.in")
	defer in.Close()
	out, _ := os.Create("445.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line, maze string
	var ok bool
	for {
		if line, ok = nextLine(s); !ok {
			line = ""
		}
		if line == "" {
			solve(out, maze)
			if !ok {
				break
			}
			maze = ""
		} else {
			if maze != "" && maze[len(maze)-1] != '!' {
				maze += "!"
			}
			maze += line
		}
	}
}
