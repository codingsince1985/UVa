// UVa 637 - Booklet Printing

package main

import (
	"fmt"
	"os"
)

var (
	out   *os.File
	pages = [2]string{"front", "back "}
)

func print(n int) {
	sheet := n / 4
	if n%4 != 0 {
		sheet++
	}
	blanks := sheet*4 - n
	left := true
	fmt.Fprintf(out, "Printing order for %d pages:\n", n)
here:
	for pageToPrint, i := 1, 1; i <= sheet; i++ {
		for _, page := range pages {
			fmt.Fprintf(out, "Sheet %d, %s: ", i, page)
			if blanks > 0 {
				if left {
					fmt.Fprintf(out, "Blank, %d\n", pageToPrint)
				} else {
					fmt.Fprintf(out, "%d, Blank\n", pageToPrint)
				}
				blanks--
			} else {
				if left {
					fmt.Fprintf(out, "%d, %d\n", n, pageToPrint)
				} else {
					fmt.Fprintf(out, "%d, %d\n", pageToPrint, n)
				}
				n--
			}
			pageToPrint++
			if pageToPrint > n {
				break here
			}
			left = !left
		}
	}
}

func main() {
	in, _ := os.Open("637.in")
	defer in.Close()
	out, _ = os.Create("637.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		print(n)
	}
}
