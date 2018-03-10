// UVa 118 - Mutant Flatworld Explorers

package main

import (
	"fmt"
	"io"
	"os"
)

var (
	dx  = []int{0, 1, 0, -1}
	dy  = []int{1, 0, -1, 0}
	dir = map[string]int{"N": 0, "E": 1, "S": 2, "W": 3}
)

func output(out io.Writer, x, y, direction int, lost bool) {
	fmt.Fprintf(out, "%d %d ", x, y)
	for i, v := range dir {
		if v == direction {
			fmt.Fprintf(out, "%s", i)
			break
		}
	}
	if lost {
		fmt.Fprint(out, " LOST")
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("118.in")
	defer in.Close()
	out, _ := os.Create("118.out")
	defer out.Close()

	var scent [][2]int
	var maxX, maxY, x, y int
	fmt.Fscanf(in, "%d%d", &maxX, &maxY)
	var orientation, instructions string
	for {
		if _, err := fmt.Fscanf(in, "%d%d%s", &x, &y, &orientation); err != nil {
			break
		}
		fmt.Fscanf(in, "%s", &instructions)
		lost := false
		direction := dir[orientation]
	here:
		for i := range instructions {
			switch instructions[i] {
			case 'R':
				direction = (direction + 1) % 4
			case 'L':
				direction = (direction + 3) % 4 // if direction is 0, direction - 1 will be -1
			case 'F':
				tmpX, tmpY := x+dx[direction], y+dy[direction]
				if tmpX < 0 || tmpX > maxX || tmpY < 0 || tmpY > maxY {
					for _, v := range scent {
						if v == [2]int{x, y} {
							continue here
						}
					}
					scent = append(scent, [2]int{x, y})
					lost = true
					break here
				}
				x, y = tmpX, tmpY
			}
		}
		output(out, x, y, direction, lost)
	}
}
