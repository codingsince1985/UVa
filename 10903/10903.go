// UVa 10903 - Rock-Paper-Scissors Tournament

package main

import (
	"fmt"
	"os"
)

type player struct{ w, l float64 }

func judge(c1, c2 string) (int, int) {
	switch {
	case c1 == "paper" && c2 == "rock" || c1 == "rock" && c2 == "scissors" || c1 == "scissors" && c2 == "paper":
		return 1, 0
	case c1 == "rock" && c2 == "paper" || c1 == "scissors" && c2 == "rock" || c1 == "paper" && c2 == "scissors":
		return 0, 1
	default:
		return 0, 0
	}
}

func output(players []player, out *os.File) {
	for _, player := range players {
		if player.w+player.l > 0 {
			fmt.Fprintf(out, "%.3f\n", player.w/(player.w+player.l))
		} else {
			fmt.Fprintln(out, "-")
		}
	}
}

func main() {
	in, _ := os.Open("10903.in")
	defer in.Close()
	out, _ := os.Create("10903.out")
	defer out.Close()

	var n, k, p1, p2 int
	var c1, c2 string
	for first := true; ; {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		players := make([]player, n)
		for fmt.Fscanf(in, "%d", &k); k > 0; k-- {
			fmt.Fscanf(in, "%d%s%d%s", &p1, &c1, &p2, &c2)
			if w, l := judge(c1, c2); w == 1 || l == 1 {
				if w == 1 {
					players[p1-1].w++
					players[p2-1].l++
				} else {
					players[p1-1].l++
					players[p2-1].w++
				}
			}
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		output(players, out)
	}
}
