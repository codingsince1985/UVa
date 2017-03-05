// UVa 10035 - Primary Arithmetic

package main

import (
	"fmt"
	"os"
)

func output(out *os.File, count int) {
	switch count {
	case 0:
		fmt.Fprintln(out, "No carry operation.")
	case 1:
		fmt.Fprintln(out, "1 carry operation.")
	default:
		fmt.Fprintf(out, "%d carry operations.\n", count)
	}
}

func main() {
	in, _ := os.Open("10035.in")
	defer in.Close()
	out, _ := os.Create("10035.out")
	defer out.Close()

	var s1, s2 string
	for {
		if fmt.Fscanf(in, "%s%s", &s1, &s2); s1 == "0" && s2 == "0" {
			break
		}
		var d1, d2, carry, count int
		l1, l2 := len(s1), len(s2)
		for l1 > 0 || l2 > 0 {
			if l1 > 0 {
				d1 = int(s1[l1-1] - '0')
				l1--
			} else {
				d1 = 0
			}
			if l2 > 0 {
				d2 = int(s2[l2-1] - '0')
				l2--
			} else {
				d2 = 0
			}
			if d1+d2+carry >= 10 {
				count++
				carry = 1
			} else {
				carry = 0
			}
		}
		output(out, count)
	}
}
