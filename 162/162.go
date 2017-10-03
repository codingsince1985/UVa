// UVa 162 - Beggar My Neighbour

package main

import (
	"fmt"
	"os"
)

var (
	values  = map[byte]int{'J': 1, 'Q': 2, 'K': 3, 'A': 4}
	table   []int
	players [][]int
)

func play(value, index int) int {
	if value == 0 {
		value = 1
	}
	for ; value > 0; value-- {
		l := len(players[index])
		if l == 0 {
			return -1
		}
		curr := players[index][l-1]
		players[index] = players[index][:l-1]
		table = append([]int{curr}, table...)
		if curr > 0 {
			return curr
		}
	}
	return 0
}

func solve() (int, int) {
	var index, value int
	for table = nil; ; {
		newValue := play(value, index)
		index = (index + 1) % 2
		if newValue == -1 {
			return 2 - index, len(players[index])
		}
		if newValue == 0 && value != 0 {
			players[index] = append(table, players[index]...)
			table = nil
		}
		value = newValue
	}
}

func main() {
	in, _ := os.Open("162.in")
	defer in.Close()
	out, _ := os.Create("162.out")
	defer out.Close()

	var c string
	for {
		if fmt.Fscanf(in, "%s", &c); c == "#" {
			break
		}
		players = make([][]int, 2)
		for i := range players {
			players[i] = make([]int, 26)
		}
		for i := range players[0] {
			if i != 0 {
				fmt.Fscanf(in, "%s", &c)
			}
			players[0][i] = values[c[1]]
			fmt.Fscanf(in, "%s", &c)
			players[1][i] = values[c[1]]
		}
		index, cards := solve()
		fmt.Fprintf(out, "%d%3d\n", index, cards)
	}
}
