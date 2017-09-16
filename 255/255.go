// UVa 255 - Correct Move

package main

import (
	"fmt"
	"os"
)

func legalAndAllowed(king, queue, move int) int {
	if queue == move {
		return -1
	}
	var low, high int
	if move < queue {
		low, high = move, queue
	} else {
		low, high = queue, move
	}
	if high-low >= 8 && (high-low)%8 == 0 {
		for i := low; i <= high; i += 8 {
			if i == king {
				return -1
			}
		}
		for i := low; i <= high; i += 8 {
			if i != queue && (i == king+1 || i == king-1 || i == king+8 || i == king-8) {
				return 0
			}
		}
		return 1
	}
	if queue/8 == move/8 {
		for i := low; i <= high; i++ {
			if i == king {
				return -1
			}
		}
		for i := low + 1; i <= high; i++ {
			if i == king+1 || i == king-1 || i == king+8 || i == king-8 {
				return 0
			}
		}
		return 1
	}
	return -1
}

func safe(move, queue int) bool {
	return move != queue+1 && move != queue-1 && move != queue-8 && move != queue+8
}

func canContinue(king, queue int) bool {
	return king%8 != 0 && safe(king-1, queue) ||
		king%8 != 7 && safe(king+1, queue) ||
		king-8 >= 0 && safe(king-8, queue) ||
		king+8 <= 63 && safe(king+8, queue)
}

func solve(king, queue, move int) string {
	if king == queue {
		return "Illegal state"
	}
	if v := legalAndAllowed(king, queue, move); v == 1 {
		if canContinue(king, move) {
			return "Continue"
		}
		return "Stop"
	} else {
		if v == -1 {
			return "Illegal move"
		}
		return "Move not allowed"
	}
}

func main() {
	in, _ := os.Open("255.in")
	defer in.Close()
	out, _ := os.Create("255.out")
	defer out.Close()

	var king, queue, move int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d", &king, &queue, &move); err != nil {
			break
		}
		fmt.Fprintln(out, solve(king, queue, move))
	}
}
