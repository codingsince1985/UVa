// UVa 10197 - Learning Portuguese

package main

import (
	"fmt"
	"os"
	"strings"
)

var pronouns = []string{
	"eu        ",
	"tu        ",
	"ele/ela   ",
	"nós       ",
	"vós       ",
	"eles/elas ",
}

func solve(v string) []string {
	if !strings.HasSuffix(v, "ar") && !strings.HasSuffix(v, "er") && !strings.HasSuffix(v, "ir") {
		return nil
	}
	root, tv := v[:len(v)-2], string(v[len(v)-2])
	conjugations := make([]string, 6)
	conjugations[0] = root + "o"
	conjugations[3] = root + tv + "mos"
	if tv == "i" {
		conjugations[1] = root + "es"
		conjugations[2] = root + "e"
		conjugations[4] = root + tv + "s"
		conjugations[5] = root + "em"
	} else {
		conjugations[1] = root + tv + "s"
		conjugations[2] = root + tv
		conjugations[4] = root + tv + "is"
		conjugations[5] = root + tv + "m"
	}
	return conjugations
}

func main() {
	in, _ := os.Open("10197.in")
	defer in.Close()
	out, _ := os.Create("10197.out")
	defer out.Close()

	var v1, v2 string
	for first := true; ; {
		if _, err := fmt.Fscanf(in, "%s%s", &v1, &v2); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "%s (to %s)\n", v1, v2)
		if conjugations := solve(v1); conjugations == nil {
			fmt.Fprintln(out, "Unknown conjugation")
		} else {
			for i, c := range conjugations {
				fmt.Fprintf(out, "%s%s\n", pronouns[i], c)
			}
		}
	}
}
