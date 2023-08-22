// UVa 443 - Humble Numbers

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, _ := os.Open("443.in")
	defer in.Close()
	out, _ := os.Create("443.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		lst := make([]int, n+1)
		lst[1] = 1
		current, p2, p3, p5, p7 := 1, 1, 1, 1, 1
		for current < n {
			for lst[p2]*2 <= lst[current] {
				p2++
			}
			for lst[p3]*3 <= lst[current] {
				p3++
			}
			for lst[p5]*5 <= lst[current] {
				p5++
			}
			for lst[p7]*7 <= lst[current] {
				p7++
			}
			current++
			lst[current] = min(lst[p2]*2, min(lst[p3]*3, min(lst[p5]*5, lst[p7]*7)))
		}

		var order string
		switch {
		case n != 11 && n%10 == 1:
			order = strconv.Itoa(n) + "st"
		case n != 12 && n%10 == 2:
			order = strconv.Itoa(n) + "nd"
		case n != 13 && n%10 == 3:
			order = strconv.Itoa(n) + "rd"
		default:
			order = strconv.Itoa(n) + "th"
		}
		fmt.Fprintf(out, "The %s humble number is %d.\n", order, lst[current])
	}
}
