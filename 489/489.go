// UVa 489 - Hangman Judge

package main

import (
	"fmt"
	"io"
	"os"
)

func solve(out io.Writer, solution, guess string) {
	chars := make(map[byte]bool)
	for i := range solution {
		chars[solution[i]] = true
	}
	count := 0
	for i := range guess {
		if chars[guess[i]] {
			delete(chars, guess[i])
			if len(chars) == 0 {
				fmt.Fprintln(out, "You win.")
				return
			}
		} else {
			if count++; count == 7 {
				fmt.Fprintln(out, "You lose.")
				return
			}
		}
	}
	fmt.Fprintln(out, "You chickened out.")
}

func main() {
	in, _ := os.Open("489.in")
	defer in.Close()
	out, _ := os.Create("489.out")
	defer out.Close()

	var kase int
	var solution, guess string
	for {
		if fmt.Fscanf(in, "%d", &kase); kase == -1 {
			break
		}
		fmt.Fprintf(out, "Round %d\n", kase)
		fmt.Fscanf(in, "%s\n%s", &solution, &guess)
		solve(out, solution, guess)
	}
}
