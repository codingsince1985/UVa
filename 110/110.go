// UVa 110 - Meta-Loopless Sorts

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	n    int
	out  *os.File
	vars []string
)

func indent(level int) { fmt.Fprint(out, strings.Repeat("  ", level)) }

func insert(level int, order []string) {
	if level == n {
		indent(level)
		fmt.Fprintf(out, "writeln(%s)\n", strings.Join(order, ","))
		return
	}
	for i := len(order); i >= 0; i-- {
		indent(level)
		if i > 0 {
			if i < len(order) {
				fmt.Fprint(out, "else ")
			}
			fmt.Fprintf(out, "if %s < %s then\n", order[i-1], vars[level])
		} else {
			fmt.Fprintln(out, "else")
		}
		insert(level+1, append(append(append([]string(nil), order[:i]...), vars[level]), order[i:]...))
	}
}

func write(n int) {
	fmt.Fprintln(out, "program sort(input,output);")
	fmt.Fprintln(out, "var")
	vars = make([]string, n)
	for i := range vars {
		vars[i] = string('a' + i)
	}
	varStr := strings.Join(vars, ",")
	fmt.Fprintf(out, "%s : integer;\n", varStr)
	fmt.Fprintln(out, "begin")
	fmt.Fprintf(out, "  readln(%s);\n", varStr)
	insert(1, vars[:1])
	fmt.Fprintln(out, "end.")
}

func main() {
	in, _ := os.Open("110.in")
	defer in.Close()
	out, _ = os.Create("110.out")
	defer out.Close()

	var m int
	fmt.Fscanf(in, "%d", &m)
	for m > 0 {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d", &n)
		write(n)
		m--
		if m > 0 {
			fmt.Fprintln(out)
		}
	}
}
