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

	var dec int
	var num string
	for {
		if fmt.Fscanf(in, "%s", &num); num[0] == '-' {
			break
		}
		if strings.HasPrefix(num, "0x") {
			dec, _ := strconv.ParseInt(num[2:], 16, 32)
			fmt.Fprintln(out, dec)
		} else {
			fmt.Sscanf(num, "%d", &dec)
			fmt.Fprintf(out, "0x%X\n", dec)
		}
	}
}
