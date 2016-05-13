// UVa 10473 - Simple Base Conversion

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("10473.in")
	defer in.Close()
	out, _ := os.Create("10473.out")
	defer out.Close()

	var num string
	for {
		if fmt.Fscanf(in, "%s", &num); num[0] == '-' {
			break
		}
		if strings.Index(num, "0x") == 0 {
			dec, _ := strconv.ParseInt(num[2:], 16, 32)
			fmt.Fprintln(out, dec)
		} else {
			dec, _ := strconv.Atoi(num)
			fmt.Fprintf(out, "0x%X\n", dec)
		}
	}
}
