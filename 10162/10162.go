// UVa 10162 - Last Digit

package main

import (
	"fmt"
	"os"
)

var lastDigits = func() []int {
	l := make([]int, 101)
	for i := 1; i <= 100; i++ {
		r := 1
		for j := 1; j <= i; j++ {
			r = (r * i) % 10
			l[i] = (l[i-1] + r) % 10
		}
	}
	return l
}()

func main() {
	in, _ := os.Open("10162.in")
	defer in.Close()
	out, _ := os.Create("10162.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%s", &line); line == "0" {
			break
		}
		var idx int
		for i := range line {
			idx = (idx*10 + int(line[i]-'0')) % 100
		}
		fmt.Fprintln(out, lastDigits[idx])
	}
}
