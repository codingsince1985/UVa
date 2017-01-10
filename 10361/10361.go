// UVa 10361 - Automatic Poetry

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in, _ := os.Open("10361.in")
	defer in.Close()
	out, _ := os.Create("10361.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n int
	for fmt.Fscanf(in, "%d", &n); n > 0 && s.Scan(); n-- {
		line := s.Text()
		idx21 := strings.Index(line, "<")
		idx22 := strings.Index(line, ">")
		idx41 := strings.LastIndex(line, "<")
		idx42 := strings.LastIndex(line, ">")

		s2 := line[idx21+1 : idx22]
		s3 := line[idx22+1 : idx41]
		s4 := line[idx41+1 : idx42]
		s5 := line[idx42+1:]
		fmt.Fprintln(out, line[:idx21]+s2+s3+s4+s5)

		s.Scan()
		line = s.Text()
		idx := strings.Index(line, "...")
		fmt.Fprintln(out, line[:idx]+s4+s3+s2+s5)
	}
}
