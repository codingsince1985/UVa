// UVa 445 - Marvelous Mazes

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var out *os.File

func nextLine(s *bufio.Scanner) (string, bool) {
	if s.Scan() {
		return s.Text(), true
	}
	return "", false
}

func outputLine(line string) {
	cnt := 0
	for i := range line {
		v := line[i]
		if v == 'b' {
			v = ' '
		}
		if v >= '0' && v <= '9' {
			cnt += int(v - '0')
		} else {
			for ; cnt > 0; cnt-- {
				fmt.Fprintf(out, "%c", v)
			}
		}
	}
	fmt.Fprintln(out)
}

func solve(maze string) {
	if maze[len(maze)-1] == '!' {
		maze = maze[:len(maze)-1]
	}
	for _, v := range strings.Split(maze, "!") {
		outputLine(v)
	}
}

func main() {
	in, _ := os.Open("445.in")
	defer in.Close()
	out, _ = os.Create("445.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line, maze string
	var ok bool
	first := true
	for {
		if line, ok = nextLine(s); !ok || line == "" {
			if first {
				first = false
			} else {
				fmt.Fprintln(out)
			}
			solve(maze)
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
