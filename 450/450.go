// UVa 450 - Little Black Book

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type faculty struct{ title, firstName, lastName, address, department, homePhone, workPhone, campusBox string }

func parse(line, department string) faculty {
	token := strings.Split(line, ",")
	return faculty{token[0], token[1], token[2], token[3], department, token[4], token[5], token[6]}
}

func output(out *os.File, faculties []faculty) {
	for _, f := range faculties {
		fmt.Fprintln(out, "----------------------------------------")
		fmt.Fprintf(out, "%s %s %s\n%s\n", f.title, f.firstName, f.lastName, f.address)
		fmt.Fprintf(out, "Department: %s\n", f.department)
		fmt.Fprintf(out, "Home Phone: %s\n", f.homePhone)
		fmt.Fprintf(out, "Work Phone: %s\n", f.workPhone)
		fmt.Fprintf(out, "Campus Box: %s\n", f.campusBox)
	}
}

func main() {
	in, _ := os.Open("450.in")
	defer in.Close()
	out, _ := os.Create("450.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n int
	var line string
	var faculties []faculty
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &n); n > 0 && s.Scan(); n-- {
		for department := s.Text(); s.Scan(); {
			if line = s.Text(); line == "" {
				break
			}
			faculties = append(faculties, parse(line, department))
		}
	}
	sort.Slice(faculties, func(i, j int) bool { return faculties[i].lastName < faculties[j].lastName })
	output(out, faculties)
}
