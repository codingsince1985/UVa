// UVa 11219 - How old are you?

package main

import (
	"fmt"
	"os"
)

type date struct{ y, m, d int }

func convert(t string) date {
	var d date
	fmt.Sscanf(t, "%d/%d/%d", &d.d, &d.m, &d.y)
	return d
}

func compare(t1, t2 string) int {
	current, birth := convert(t1), convert(t2)
	age := current.y - birth.y
	if current.m < birth.m || current.m == birth.m && current.d < birth.d {
		age--
	}
	return age
}

func main() {
	in, _ := os.Open("11219.in")
	defer in.Close()
	out, _ := os.Create("11219.out")
	defer out.Close()

	var kase int
	fmt.Fscanf(in, "%d", &kase)
	var t1, t2 string
	for i := 1; i <= kase; i++ {
		fmt.Fscanf(in, "\n%s", &t1)
		fmt.Fscanf(in, "%s", &t2)
		years := compare(t1, t2)
		fmt.Fprintf(out, "Case #%d: ", i)
		switch {
		case years < 0:
			fmt.Fprintln(out, "Invalid birth date")
		case years > 130:
			fmt.Fprintln(out, "Check birth date")
		default:
			fmt.Fprintf(out, "%d\n", years)
		}
	}
}
