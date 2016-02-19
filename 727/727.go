// UVa 727 - Equation

package main

import (
	"fmt"
	"os"
)

const MAX = 50

var count int
var stack [MAX]string
var out *os.File

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
	fmt.Fscanf(in, "%d\n\n", &n)
	var s string
	for i := 0; i < n; i++ {
		for {
			s = ""
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
