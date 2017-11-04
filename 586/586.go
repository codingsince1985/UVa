// UVa 586 - Instant Complexity

package main

import (
	"fmt"
	"os"
)

const max = 10

var (
	in           *os.File
	coefficients []int
)

func dfs(n, k int) {
	var coefficient, tmp int
	var command string
	for {
		if fmt.Fscanf(in, "%s", &command); command == "END" {
			coefficients[n] += coefficient * k
			return
		}
		switch command {
		case "LOOP":
			if fmt.Fscanf(in, "%s", &command); command == "n" {
				dfs(n+1, k)
			} else {
				fmt.Sscanf(command, "%d", &tmp)
				dfs(n, k*tmp)
			}
		case "OP":
			fmt.Fscanf(in, "%d", &tmp)
			coefficient += tmp
		}
	}
}

func output(out *os.File, kase int) {
	fmt.Fprintf(out, "Program #%d\n", kase)
	fmt.Fprint(out, "Runtime = ")
	var started bool
	for power := max; power >= 0; power-- {
		if coefficients[power] != 0 {
			if started {
				fmt.Fprint(out, "+")
			} else {
				started = true
			}
			if power == 0 {
				fmt.Fprint(out, coefficients[power])
			} else {
				if coefficients[power] > 1 {
					fmt.Fprintf(out, "%d*", coefficients[power])
				}
				fmt.Fprint(out, "n")
				if power > 1 {
					fmt.Fprintf(out, "^%d", power)
				}
			}
		}
	}
	if !started {
		fmt.Fprint(out, 0)
	}
	fmt.Fprintln(out)
	fmt.Fprintln(out)
}

func main() {
	in, _ = os.Open("586.in")
	defer in.Close()
	out, _ := os.Create("586.out")
	defer out.Close()

	var tmp string
	var k int
	fmt.Fscanf(in, "%d", &k)
	for kase := 1; kase <= k; kase++ {
		fmt.Fscanf(in, "%s", &tmp)
		coefficients = make([]int, max+1)
		dfs(0, 1)
		output(out, kase)
		fmt.Fscanln(in)
	}
}
