// UVa 740 - Baudot Data Communication Code

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("740.in")
	defer in.Close()
	out, _ := os.Create("740.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	downs := s.Text()
	s.Scan()
	ups := s.Text()

	for s.Scan() {
		line := s.Text()
		for i, down := 0, true; i+5 <= len(line); i += 5 {
			var idx byte
			for j := 0; j < 5; j++ {
				idx = idx*2 + line[i+j] - '0'
			}
			switch {
			case down && idx == 31 || !down && idx == 27:
				down = !down
			case down:
				fmt.Fprintf(out, "%c", downs[idx])
			default:
				fmt.Fprintf(out, "%c", ups[idx])
			}
		}
		fmt.Fprintln(out)
	}
}
