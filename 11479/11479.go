// UVa 11479 - Is this the easiest problem?

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11479.in")
	defer in.Close()
	out, _ := os.Create("11479.out")
	defer out.Close()

	var t, s1, s2, s3 int
	fmt.Fscanf(in, "%d", &t)
	for i := 1; i <= t; i++ {
		if _, err := fmt.Fscanf(in, "%d%d%d", &s1, &s2, &s3); err != nil {
			break
		}
		fmt.Fprintf(out, "Case %d: ", i)
		switch {
		case s1 <= 0 || s2 <= 0 || s3 <= 0 || s1+s2 <= s3 || s2+s3 <= s1 || s1+s3 <= s2:
			fmt.Fprintln(out, "Invalid")
		case s1 == s2 && s1 == s3:
			fmt.Fprintln(out, "Equilateral")
		case s1 == s2 || s1 == s3 || s2 == s3:
			fmt.Fprintln(out, "Isosceles")
		default:
			fmt.Fprintln(out, "Scalene")
		}
	}
}
