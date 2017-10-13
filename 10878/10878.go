// UVa 10878 - Decode the tape

package main

import (
	"bufio"
	"fmt"
	"os"
)

func decode(line string) byte {
	line = line[2:6] + line[7:10]
	var ascii byte
	for i := range line {
		ascii *= 2
		if line[i] == 'o' {
			ascii += 1
		}
	}
	return ascii
}

func main() {
	in, _ := os.Open("10878.in")
	defer in.Close()
	out, _ := os.Create("10878.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	for s.Scan() {
		if line := s.Text(); line != "___________" {
			fmt.Fprintf(out, "%c", decode(line))
		}
	}
}
