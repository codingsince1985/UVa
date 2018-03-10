// UVa 448 - OOPS!

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type opcode struct {
	op       string
	operands int
}

var (
	instruction = []opcode{
		{"ADD", 2}, {"SUB", 2}, {"MUL", 2}, {"DIV", 2}, {"MOV", 2}, {"BREQ", 1}, {"BRLE", 1}, {"BRLS", 1},
		{"BRGE", 1}, {"BRGR", 1}, {"BRNE", 1}, {"BR", 1}, {"AND", 3}, {"OR", 3}, {"XOR", 3}, {"NOT", 1},
	}
	operandPrefix = []string{"R", "$", "PC+", ""}
)

func decode(out io.Writer, code string) {
	var num1, num2 uint16
	for r := strings.NewReader(code); ; {
		if _, err := fmt.Fscanf(r, "%1x", &num1); err != nil {
			break
		}
		fmt.Fprintf(out, "%s ", instruction[num1].op)
		for i := 0; i < instruction[num1].operands; i++ {
			if i > 0 {
				fmt.Fprint(out, ",")
			}
			fmt.Fscanf(r, "%4x", &num2)
			fmt.Fprintf(out, "%s%d", operandPrefix[num2>>14], num2<<2>>2)
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in, _ := os.Open("448.in")
	defer in.Close()
	out, _ := os.Create("448.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var code strings.Builder
	for s.Scan() {
		code.WriteString(s.Text())
	}
	decode(out, code.String())
}
