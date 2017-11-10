// UVa 168 - Theseus and the Minotaur

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	maze      map[byte][]byte
	k         int
	candles   []string
	candleMap map[byte]bool
	trapped   byte
)

func buildMaze(token string) map[byte][]byte {
	maze = make(map[byte][]byte)
	for _, path := range strings.Split(token, ";") {
		t := strings.Split(path, ":")
		maze[t[0][0]] = []byte(t[1])
	}
	return maze
}

func dfs(m, t byte, step int) bool {
	if step == k {
		candleMap[m] = true
		candles = append(candles, string(m))
		step = 0
	}
	done := true
	for _, node := range maze[m] {
		if node != t && !candleMap[node] {
			if dfs(node, m, step+1) {
				return true
			}
			done = false
		}
	}
	if done {
		trapped = m
	}
	return done
}

func solve(minotaur, theseus byte) {
	candles = nil
	candleMap = make(map[byte]bool)
	dfs(minotaur, theseus, 1)
}

func main() {
	in, _ := os.Open("168.in")
	defer in.Close()
	out, _ := os.Create("168.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	var minotaur, theseus byte
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		tokens := strings.Split(line, ". ")
		buildMaze(tokens[0])
		fmt.Sscanf(tokens[1], "%c %c %d", &minotaur, &theseus, &k)
		solve(minotaur, theseus)
		fmt.Fprintf(out, "%s /%c\n", strings.Join(candles, " "), trapped)
	}
}
