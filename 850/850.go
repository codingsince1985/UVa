// UVa 850 - Crypt Kicker II

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const base = "the quick brown fox jumps over the lazy dog"

func match(line string) map[byte]byte {
	if len(line) != len(base) {
		return nil
	}
	dict := make(map[byte]byte)
	for i := range line {
		if b, ok := dict[line[i]]; ok && b != base[i] {
			return nil
		}
		dict[line[i]] = base[i]
	}
	return dict
}

func decodeLine(line string, dict map[byte]byte) string {
	d := make([]byte, len(line))
	for i := range line {
		d[i] = dict[line[i]]
	}
	return string(d)
}

func decode(lines []string) []string {
	var dict map[byte]byte
	for _, line := range lines {
		if dict = match(line); dict != nil {
			break
		}
	}
	if dict == nil {
		return []string{"no solution"}
	}
	decoded := make([]string, len(lines))
	for i, line := range lines {
		decoded[i] = decodeLine(line, dict)
	}
	return decoded
}

func main() {
	in, _ := os.Open("850.in")
	defer in.Close()
	out, _ := os.Create("850.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	var line string
	for ; kase > 0; kase-- {
		var lines []string
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			lines = append(lines, line)
		}
		fmt.Fprintln(out, strings.Join(decode(lines), "\n"))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
