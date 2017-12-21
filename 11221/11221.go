// UVa 11221 - Magic square palindromes.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var keepLetter = func(r rune) rune {
	if r >= rune('a') && r <= rune('z') {
		return r
	}
	return -1
}

func isPalindromes(l int, line string) bool {
	lastIndex := l - 1
	for i := 0; i < lastIndex/2; i++ {
		if line[i] != line[lastIndex-i] {
			return false
		}
	}
	return true
}

func reShape(side int, line string) string {
	newLine := make([]byte, side*side)
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			newLine[i*side+j] = line[i+j*side]
		}
	}
	return string(newLine)
}

func solve(line string) int {
	newLine := strings.Map(keepLetter, line)
	length := len(newLine)
	if side := int(math.Sqrt(float64(length)) + .5); side*side == length && isPalindromes(length, newLine) && isPalindromes(length, reShape(side, newLine)) {
		return side
	}
	return -1
}

func main() {
	in, _ := os.Open("11221.in")
	defer in.Close()
	out, _ := os.Create("11221.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	for i := 1; i <= kase && s.Scan(); i++ {
		fmt.Fprintf(out, "Case #%d:\n", i)
		if side := solve(s.Text()); side == -1 {
			fmt.Fprintln(out, "No magic :(")
		} else {
			fmt.Fprintln(out, side)
		}
	}
}
