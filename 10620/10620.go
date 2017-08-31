// UVa 10620 - A Flea on a Chessboard

package main

import (
	"fmt"
	"os"
)

func white(s, x, y int) bool { return x%s != 0 && y%s != 0 && (x/s)%2+(y/s)%2 == 1 }

func solve(s, x, y, dx, dy int) string {
	for i := 0; i < 2*s; i, x, y = i+1, x+dx, y+dy {
		if white(s, x, y) {
			return fmt.Sprintf("After %d jumps the flea lands at (%d, %d).", i, x, y)
		}
	}
	return "The flea cannot escape from black squares."
}

func main() {
	in, _ := os.Open("10620.in")
	defer in.Close()
	out, _ := os.Create("10620.out")
	defer out.Close()

	var s, x, y, dx, dy int
	for {
		if fmt.Fscanf(in, "%d%d%d%d%d", &s, &x, &y, &dx, &dy); s == 0 {
			break
		}
		fmt.Fprintln(out, solve(s, x, y, dx, dy))
	}
}
