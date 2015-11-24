// UVa 727 - Equation

package main

import (
	"fmt"
	"os"
)

const MAX = 50

var count int
var stack[MAX] string
var out *os.File

func pop() {
	count --
	fmt.Fprintf(out, "%s", stack[count])
}

func push(s string) {
	stack[count] = s
	count ++
}

func main() {
	in, _ := os.Open("727.in")
	out, _ = os.Create("727.out")
	var n int
	fmt.Fscanf(in, "%d\n\n", &n)
	var s string
	for i := 0; i < n; i++ {
		for {
			s = ""
			fmt.Fscanf(in, "%s\n", &s)
			if len(s) == 0 {
				break
			}
			if s == "+" || s == "-" {
				for count > 0 && stack[count - 1] != "(" {
					pop();
				}
				push(s);
			} else if s == "*" || s == "/" {
				for count > 0 && stack[count - 1] != "(" && stack[count - 1] != "+" && stack[count - 1] != "-" {
					pop()
				}
				push(s)
			} else if s == "(" {
				push(s)
			} else if s == ")" {
				for stack[count - 1] != "(" {
					pop()
				}
				count--
			} else {
				fmt.Fprint(out, s)
			}

		}
		for count != 0 {
			pop()
		}
		fmt.Fprintln(out)
	}
	in.Close()
	out.Close()
}
