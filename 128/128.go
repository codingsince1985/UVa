// UVa 128 - Software CRC

package main

import (
	"bufio"
	"fmt"
	"os"
)

const g = 34943

func main() {
	in, _ := os.Open("128.in")
	defer in.Close()
	out, _ := os.Create("128.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if line := s.Text(); line != "#" {
			num := 0
			for i := range line {
				num = (num<<8 + int(line[i])) % g
			}
			num = (g - (num << 16 % g)) % g
			fmt.Fprintf(out, "%02X %02X\n", num>>8, num%256)
		}
	}
}
