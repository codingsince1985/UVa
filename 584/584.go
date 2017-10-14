// UVa 584 - Bowling

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const plays = 10

type frame struct {
	scores        []int
	strike, spare bool
}

func calc(frames []frame) int {
	var score int
	for i := 0; i < plays; i++ {
		score += frames[i].scores[0]
		if !frames[i].strike {
			score += frames[i].scores[1]
		}
		if frames[i].spare || frames[i].strike {
			score += frames[i+1].scores[0]
			if frames[i].strike {
				if frames[i+1].strike {
					score += frames[i+2].scores[0]
				} else {
					score += frames[i+1].scores[1]
				}
			}
		}
	}
	return score
}

func solve(line string) int {
	var idx, s int
	frames := make([]frame, plays+2)
	for _, score := range strings.Split(line, " ") {
		switch score {
		case "X":
			frames[idx].scores = []int{10}
			frames[idx].strike = true
			idx++
		case "/":
			frames[idx].scores = append(frames[idx].scores, 10-frames[idx].scores[0])
			frames[idx].spare = true
			idx++
		default:
			fmt.Sscanf(score, "%d", &s)
			if frames[idx].scores = append(frames[idx].scores, s); len(frames[idx].scores) == 2 {
				idx++
			}
		}
	}
	return calc(frames)
}

func main() {
	in, _ := os.Open("584.in")
	defer in.Close()
	out, _ := os.Create("584.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "Game Over" {
			break
		}
		fmt.Fprintln(out, solve(line))
	}
}
