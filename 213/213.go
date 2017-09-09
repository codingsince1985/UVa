// UVa 213 - Message Decoding

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s *bufio.Scanner

func buildKey(header string) map[string]byte {
	keyMap := make(map[string]byte)
	var key int64
	digit := uint(1)
	for i := range header {
		if key == (2<<(digit-1) - 1) {
			key = 0
			digit++
		}
		binary := strconv.FormatInt(key, 2)
		binary = strings.Repeat("0", int(digit)-len(binary)) + binary
		keyMap[binary] = header[i]
		key++
	}
	return keyMap
}

func getLength(l string) int {
	var n int
	fmt.Sscanf(l, "%b", &n)
	return n
}

func solve(out *os.File, header string) {
	keyMap := buildKey(header)
	var l int
	var line, code string
	s.Scan()
	for line = s.Text(); line != "000"; {
		l, line = getLength(line[:3]), line[3:]
		for {
			code, line = line[:l], line[l:]
			if len(line) < l {
				s.Scan()
				line += s.Text()
			}
			if strings.Count(code, "1") == l {
				break
			}
			fmt.Fprintf(out, "%c", keyMap[code])
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("213.in")
	defer in.Close()
	out, _ := os.Create("213.out")
	defer out.Close()

	s = bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "" {
			break
		}
		solve(out, line)
	}
}
