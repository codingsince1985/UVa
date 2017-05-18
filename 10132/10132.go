// UVa 10132 - File Fragmentation

package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	length int
	lines  []string
)

func validate(file string) bool {
	fileLength := len(file)
	visited := make([]bool, length)
here:
	for i := 0; i < length/2; i++ {
		expectedLength := fileLength - len(lines[i])
		for j := length / 2; j < length; j++ {
			if !visited[j] && len(lines[j]) == expectedLength {
				if file == lines[i]+lines[j] || file == lines[j]+lines[i] {
					visited[i], visited[j] = true, true
					continue here
				}
			}
		}
		return false
	}
	for i := range visited {
		if !visited[i] {
			return false
		}
	}
	return true
}

func solve() string {
	sort.Slice(lines, func(i, j int) bool { return len(lines[i]) < len(lines[j]) })
	min := len(lines[0])
	length = len(lines)
	for i, line := range lines {
		if len(line) > min {
			break
		}
		if file := lines[0] + lines[length-i-1]; validate(file) {
			return file
		}
		if file := lines[length-i-1] + lines[0]; validate(file) {
			return file
		}
	}
	return ""
}

func main() {
	in, _ := os.Open("10132.in")
	defer in.Close()
	out, _ := os.Create("10132.out")
	defer out.Close()

	var kase int
	first := true
	for fmt.Fscanf(in, "%d\n\n", &kase); kase > 0; kase-- {
		lines = nil
		for {
			var line string
			if fmt.Fscanf(in, "%s", &line); line == "" {
				break
			}
			lines = append(lines, line)
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, solve())
	}
}
