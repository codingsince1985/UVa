// UVa 727 - Equation

package main

import (
	"fmt"
	"os"
)

const max = 50

var (
	count int
	stack [max]string
	out   *os.File
)

func pop() {
	count--
	fmt.Fprintf(out, "%s", stack[count])
}

func push(s string) {
	stack[count] = s
	count++
}

func main() {
	in, _ := os.Open("727.in")
	defer in.Close()
	out, _ = os.Create("727.out")
	defer out.Close()

	var n int
	for fmt.Fscanf(in, "%d\n\n", &n); n > 0; n-- {
		for {
			var s string
			if fmt.Fscanf(in, "%s\n", &s); len(s) == 0 {
				break
			}
			switch s {
			case "+", "-":
				for count > 0 && stack[count-1] != "(" {
					pop()
				}
				push(s)
			case "*", "/":
				for count > 0 && stack[count-1] != "(" && stack[count-1] != "+" && stack[count-1] != "-" {
					pop()
				}
				push(s)
			case "(":
				push(s)
			case ")":
				for stack[count-1] != "(" {
					pop()
				}
				count--
			default:
				fmt.Fprint(out, s)
			}
		}
		for count != 0 {
			pop()
		}
		fmt.Fprintln(out)
	}
}
