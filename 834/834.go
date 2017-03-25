// UVa 834 - Continued Fractions

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in, _ := os.Open("834.in")
	defer in.Close()
	out, _ := os.Create("834.out")
	defer out.Close()

	var a, b int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &a, &b); err != nil {
			break
		}
		lst := []int{a / b}
		a %= b
		for a != 1 {
			lst = append(lst, b/a)
			b %= a
			a, b = b, a
		}
		lst = append(lst, b)
		str := fmt.Sprint(lst)
		str = strings.Replace(str, " ", ";", 1)
		str = strings.Replace(str, " ", ",", -1)
		fmt.Fprintln(out, str)
	}
}
