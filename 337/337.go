// UVa 337 - Interpreting Control Sequences

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type terminal struct {
	text        [10][10]byte
	overwrite   bool
	row, column int
}

var term terminal

func clear() {
	for i := range term.text {
		for j := range term.text[i] {
			term.text[i][j] = ' '
		}
	}
}

func initialize() {
	var text [10][10]byte
	term = terminal{text, true, 0, 0}
	clear()
}

func printByte(c byte) {
	if !term.overwrite {
		for i := 9; i > term.column; i-- {
			term.text[term.row][i] = term.text[term.row][i-1]
		}
	}
	term.text[term.row][term.column] = c
	if term.column < 9 {
		term.column++
	}
}

func display(input string) {
	initialize()
	for i := 0; i < len(input); i++ {
		if input[i] != '^' {
			printByte(input[i])
			continue
		}
		i++
		switch {
		case input[i] >= '0' && input[i] <= '9':
			term.row = int(input[i] - '0')
			i++
			term.column = int(input[i] - '0')
		case input[i] == 'b':
			term.column = 0
		case input[i] == 'c':
			clear()
		case input[i] == 'd':
			if term.row < 9 {
				term.row++
			}
		case input[i] == 'e':
			for i := term.column; i <= 9; i++ {
				term.text[term.row][i] = ' '
			}
		case input[i] == 'h':
			term.row, term.column = 0, 0
		case input[i] == 'i':
			term.overwrite = false
		case input[i] == 'l':
			if term.column > 0 {
				term.column--
			}
		case input[i] == 'o':
			term.overwrite = true
		case input[i] == 'r':
			if term.column < 9 {
				term.column++
			}
		case input[i] == 'u':
			if term.row > 0 {
				term.row--
			}
		case input[i] == '^':
			printByte('^')
		}
	}
}

func output(out io.Writer, kase int) {
	fmt.Fprintf(out, "Case %d\n+----------+\n", kase)
	for _, row := range term.text {
		fmt.Fprint(out, "|")
		fmt.Fprintln(out, string(row[:])+"|")
	}
	fmt.Fprintln(out, "+----------+")
}

func main() {
	in, _ := os.Open("337.in")
	defer in.Close()
	out, _ := os.Create("337.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n int
	for kase := 1; s.Scan(); kase++ {
		var inputBuilder strings.Builder
		for fmt.Sscanf(s.Text(), "%d", &n); n > 0 && s.Scan(); n-- {
			inputBuilder.WriteString(s.Text())
		}
		input := inputBuilder.String()
		if input == "" {
			break
		}
		display(input)
		output(out, kase)
	}
}
