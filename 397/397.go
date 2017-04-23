// UVa 397 - Equation Elation

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var out *os.File

func parse(line string) []string {
	var stack []string
	var num int
	var ch byte
	var variable string
	r := strings.NewReader(line)
	for {
		fmt.Fscanf(r, "%d", &num)
		stack = append(stack, strconv.Itoa(num))
		for {
			if fmt.Fscanf(r, "%c", &ch); ch != ' ' {
				break
			}
		}
		stack = append(stack, string(ch))
		if ch == '=' {
			break
		}
	}
	fmt.Fscanf(r, "%s", &variable)
	stack = append(stack, variable)
	return stack
}

func operations(stack []string) (int, int) {
	var md, as int
	for _, element := range stack {
		switch element {
		case "*", "/":
			md++
		case "+", "-":
			as++
		}
	}
	return md, as
}

func pass(times int, ops [2]string, stack []string) []string {
	var num1, num2 int
	for idx := 0; times > 0; idx++ {
		num1, _ = strconv.Atoi(stack[idx])
		if idx++; stack[idx] == ops[0] || stack[idx] == ops[1] {
			num2, _ = strconv.Atoi(stack[idx+1])
			switch stack[idx] {
			case "*":
				stack[idx-1] = strconv.Itoa(num1 * num2)
			case "/":
				stack[idx-1] = strconv.Itoa(num1 / num2)
			case "+":
				stack[idx-1] = strconv.Itoa(num1 + num2)
			case "-":
				stack[idx-1] = strconv.Itoa(num1 - num2)
			}
			stack = append(stack[:idx], stack[idx+2:]...)
			fmt.Fprintf(out, "%s\n", strings.Join(stack, " "))
			idx -= 2
			times--
		}
	}
	return stack
}

func solve(line string) {
	stack := parse(line)
	fmt.Fprintf(out, "%s\n", strings.Join(stack, " "))
	md, as := operations(stack)
	stack = pass(md, [2]string{"*", "/"}, stack)
	stack = pass(as, [2]string{"+", "-"}, stack)
}

func main() {
	in, _ := os.Open("397.in")
	defer in.Close()
	out, _ = os.Create("397.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	first := true
	for s.Scan() {
		if line = s.Text(); line == "" {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(line)
	}
}
