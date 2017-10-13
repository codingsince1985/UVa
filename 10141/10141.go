// UVa 10141 - Request for Proposal

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10141.in")
	defer in.Close()
	out, _ := os.Create("10141.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n, p, kase int
	for s.Scan() {
		if fmt.Sscanf(s.Text(), "%d%d", &n, &p); n == 0 && p == 0 {
			break
		}
		for i := 0; i < n && s.Scan(); i++ {
		}

		lowest := math.MaxFloat64
		var name, choice string
		var max, met int
		var price float64
		for i := 0; i < p && s.Scan(); i++ {
			name = s.Text()
			s.Scan()
			fmt.Sscanf(s.Text(), "%f%d", &price, &met)
			for j := 0; j < met && s.Scan(); j++ {
			}
			if met > max || met == max && price < lowest {
				choice, max, lowest = name, met, price
			}
		}
		kase++
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "RFP #%d\n%s\n", kase, choice)
	}
}
