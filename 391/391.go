// UVa 391 - Mark-up

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(lines []string) []string {
	bold, italics, fontSize, markUp := false, false, false, true
	var newLines []string
	var ch byte
	var f float64
	for _, line := range lines {
		var newLine strings.Builder
		for r := strings.NewReader(line); ; {
			if _, err := fmt.Fscanf(r, "%c", &ch); err != nil {
				break
			}
			if ch != '\\' {
				newLine.WriteByte(ch)
				continue
			}
			fmt.Fscanf(r, "%c", &ch)
			if markUp {
				switch ch {
				case 'b':
					bold = !bold
				case 'i':
					italics = !italics
				case 's':
					if fontSize = !fontSize; fontSize {
						fmt.Fscanf(r, "%f", &f)
					}
				case '*':
					markUp = false
				default:
					newLine.WriteByte(ch)
				}
			} else {
				if ch == '*' {
					markUp = true
				} else {
					newLine.WriteString("\\")
					newLine.WriteByte(ch)
				}
			}
		}
		newLines = append(newLines, newLine.String())
	}
	return newLines
}

func main() {
	in, _ := os.Open("391.in")
	defer in.Close()
	out, _ := os.Create("391.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	fmt.Fprintln(out, strings.Join(solve(lines), "\n"))
}
