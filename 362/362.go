// UVa 362 - 18,000 Seconds Remaining

package main

import (
	"fmt"
	"io"
	"os"
)

func output(out io.Writer, size, b5 int) {
	if b5 == 0 {
		fmt.Fprintln(out, "   Time remaining: stalled")
	} else {
		remaining := float64(size*5) / float64(b5)
		if remaining != float64(int(remaining)) {
			remaining = float64(int(remaining + 1))
		}
		fmt.Fprintf(out, "   Time remaining: %.0f seconds\n", remaining)
	}
}

func main() {
	in, _ := os.Open("362.in")
	defer in.Close()
	out, _ := os.Create("362.out")
	defer out.Close()

	var size, b int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &size); size == 0 {
			break
		}
		fmt.Fprintf(out, "Output for data set %d, %d bytes:\n", kase, size)
		var count int
		for soFar, b5 := 0, 0; soFar < size; {
			fmt.Fscanf(in, "%d", &b)
			b5 += b
			soFar += b
			if count++; count%5 == 0 {
				output(out, size-soFar, b5)
				b5 = 0
			}
		}
		fmt.Fprintf(out, "Total time: %d seconds\n\n", count)
	}
}
