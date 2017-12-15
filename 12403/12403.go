// UVa 12403 - Save Setu

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("12403.in")
	defer in.Close()
	out, _ := os.Create("12403.out")
	defer out.Close()

	var t, amount, sum int
	var operation string
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		switch fmt.Fscanf(in, "%s", &operation); operation {
		case "donate":
			fmt.Fscanf(in, "%d", &amount)
			sum += amount
		case "report":
			fmt.Fprintln(out, sum)
		}
	}
}
