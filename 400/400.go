// UVa 400 - Unix ls

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

const max = 60

func output(out io.Writer, files []string, cols, maxLen int) {
	for i := 0; i < max; i++ {
		fmt.Fprint(out, "-")
	}
	fmt.Fprintln(out)
	rows := len(files) / cols
	if len(files)%rows != 0 {
		rows++
	}
	format := fmt.Sprintf("%%-%ds", maxLen+2)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols-1; j++ {
			idx := i + j*rows
			if idx < len(files) {
				fmt.Fprintf(out, format, files[idx])
			}
		}
		idx := i + (cols-1)*rows
		if idx < len(files) {
			fmt.Fprintln(out, files[idx])
		}
	}
	if rows*cols != len(files) {
		fmt.Fprintln(out)
	}
}

func main() {
	in, _ := os.Open("400.in")
	defer in.Close()
	out, _ := os.Create("400.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		files := make([]string, n)
		maxLen := 0
		for i := range files {
			fmt.Fscanf(in, "%s", &files[i])
			length := len(files[i])
			if length > maxLen {
				maxLen = length
			}
		}
		cols := (max + 2) / (maxLen + 2)
		sort.Strings(files)
		output(out, files, cols, maxLen)
	}
}
