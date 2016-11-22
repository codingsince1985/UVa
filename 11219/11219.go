// UVa 11219 - How old are you?

package main

import (
	"fmt"
	"os"
)

type date struct{ y, m, d int }

func convert(t string) date {
	var y, m, d int
	fmt.Sscanf(t, "%d/%d/%d", &d, &m, &y)
	return date{y, m, d}
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
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%s", &t1)
		fmt.Fscanf(in, "%s", &t2)
		years := compare(t1, t2)
		switch {
		case years < 0:
			fmt.Fprintf(out, "Case #%d: Invalid birth date\n", i)
		case years > 130:
			fmt.Fprintf(out, "Case #%d: Check birth date\n", i)
		default:
			fmt.Fprintf(out, "Case #%d: %d\n", i, years)
		}
	}
}
