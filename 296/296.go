// UVa 296 - Safebreaker

package main

import (
	"fmt"
	"os"
)

type guess struct{ code, correct, misplace int }

func digits(code int) []int {
	return []int{code / 1000, (code / 100) % 10, (code / 10) % 10, code % 10}
}

func compare(c1, c2 int) (int, int) {
	var correct, misplace int
	digits1, digits2 := digits(c1), digits(c2)
	done1, done2 := make(map[int]bool), make(map[int]bool)
	for i, d := range digits1 {
		if digits2[i] == d {
			correct++
			done1[i], done2[i] = true, true
		}
	}
	for i, d1 := range digits1 {
		if !done1[i] {
			for j, d2 := range digits2 {
				if !done2[j] && d1 == d2 {
					misplace++
					done1[i], done2[j] = true, true
				}
			}
		}
	}
	return correct, misplace
}

func match(code int, guesses []guess) bool {
	for _, g := range guesses {
		if correct, misplace := compare(code, g.code); correct != g.correct || misplace != g.misplace {
			return false
		}
	}
	return true
}

func solve(guesses []guess) (int, int) {
	var count, code int
	for i := 0; i <= 9999; i++ {
		if match(i, guesses) {
			count++
			code = i
		}
	}
	return count, code
}

func main() {
	in, _ := os.Open("296.in")
	defer in.Close()
	out, _ := os.Create("296.out")
	defer out.Close()

	var n, g int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &g)
		guesses := make([]guess, g)
		for i := range guesses {
			fmt.Fscanf(in, "%d%d/%d", &guesses[i].code, &guesses[i].correct, &guesses[i].misplace)
		}
		switch matched, code := solve(guesses); matched {
		case 1:
			fmt.Fprintln(out, code)
		case 0:
			fmt.Fprintln(out, "impossible")
		default:
			fmt.Fprintln(out, "indeterminate")
		}
	}
}
