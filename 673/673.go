// UVa 673 - Parentheses Balance

package main

import (
	"bufio"
	"fmt"
	"os"
)

func match(a, b byte) bool { return a == ')' && b == '(' || a == ']' && b == '[' }

func valid(line string) bool {
	chars := []byte(line)
	var stack []byte
	for _, v := range chars {
		if v != '(' && v != ')' && v != '[' && v != ']' {
			continue
		}
		if v == '(' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 && match(v, stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	in, _ := os.Open("673.in")
	defer in.Close()
	out, _ := os.Create("673.out")
	defer out.Close()

	var n int
	var line string
	fmt.Fscanf(in, "%d", &n)

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for n > 0 {
		scanner.Scan()
		if line = scanner.Text(); valid(line) {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
		n--
	}
}
