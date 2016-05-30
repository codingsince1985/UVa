// UVa 621 - Secret Research

package main

import (
	"fmt"
	"os"
)

func solve(str string) string {
	switch {
	case str == "1" || str == "4" || str == "78":
		return "+"
	case str[len(str)-2:] == "35":
		return "-"
	case str[0] == '9' && str[len(str)-1] == '4':
		return "*"
	case str[:3] == "190":
		return "?"
	default:
		return ""
	}
}

func main() {
	in, _ := os.Open("621.in")
	defer in.Close()
	out, _ := os.Create("621.out")
	defer out.Close()

	var n int
	var str string
	fmt.Fscanf(in, "%d", &n)
	for n > 0 {
		fmt.Fscanf(in, "%s", &str)
		fmt.Fprintln(out, solve(str))
		n--
	}
}
