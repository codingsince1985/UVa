// UVa 450 - Little Black Book

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	faculty   struct{ title, firstName, lastName, address, department, homePhone, workPhone, campusBox string }
	faculties []faculty
)

func (f faculties) Len() int { return len(f) }

func (f faculties) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func (f faculties) Less(i, j int) bool { return f[i].lastName < f[j].lastName }

func parse(line, department string) faculty {
	token := strings.Split(line, ",")
	return faculty{token[0], token[1], token[2], token[3], department, token[4], token[5], token[6]}
}

func output(out *os.File, fs faculties) {
	for _, f := range fs {
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

	var line string
	var fs faculties
	s.Scan()
	for n, _ := strconv.Atoi(s.Text()); n > 0 && s.Scan(); n-- {
		for department := s.Text(); s.Scan(); {
			if line = s.Text(); line == "" {
				break
			}
			fs = append(fs, parse(line, department))
		}
	}
	sort.Sort(fs)
	output(out, fs)
}
