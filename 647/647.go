// UVa 647 - Chutes and Ladders

package main

import (
	"fmt"
	"os"
)

func solve(throws []int, players int, ladders map[int]int, loseTurns, extraTurns map[int]bool) int {
	squares := make([]int, players)
	loseTurn := make(map[int]bool)
	for i, player := 0, 0; i < len(throws); {
		if loseTurn[player] {
			loseTurn[player] = false
			player = (player + 1) % players
			continue
		}
		if squares[player] >= 95 && squares[player]+throws[i] > 100 {
			player = (player + 1) % players
			i++
			continue
		}
		if squares[player] += throws[i]; squares[player] == 100 {
			return player + 1
		}
		i++
		if s, ok := ladders[squares[player]]; ok {
			squares[player] = s
		}
		if loseTurns[squares[player]] {
			loseTurn[player] = true
		}
		if !extraTurns[squares[player]] {
			player = (player + 1) % players
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("647.in")
	defer in.Close()
	out, _ := os.Create("647.out")
	defer out.Close()

	var tmp, players, s1, s2 int
	var throws []int
	for {
		if fmt.Fscanf(in, "%d", &tmp); tmp == 0 {
			break
		}
		throws = append(throws, tmp)
	}
	for {
		if fmt.Fscanf(in, "%d", &players); players == 0 {
			break
		}
		ladders := make(map[int]int)
		for {
			if fmt.Fscanf(in, "%d%d", &s1, &s2); s1 == 0 && s2 == 0 {
				break
			}
			ladders[s1] = s2
		}
		loseTurns, extraTurns := make(map[int]bool), make(map[int]bool)
		for {
			if fmt.Fscanf(in, "%d", &tmp); tmp == 0 {
				break
			}
			if tmp < 0 {
				loseTurns[-tmp] = true
			} else {
				extraTurns[tmp] = true
			}
		}
		fmt.Fprintln(out, solve(throws, players, ladders, loseTurns, extraTurns))
	}
}
