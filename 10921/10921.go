// UVa 10921 - Find the Telephone

package main

import (
	"fmt"
	"os"
	"strings"
)

var dict = map[byte]byte{
	'A': '2',
	'B': '2',
	'C': '2',
	'D': '3',
	'E': '3',
	'F': '3',
	'G': '4',
	'H': '4',
	'I': '4',
	'J': '5',
	'K': '5',
	'L': '5',
	'M': '6',
	'N': '6',
	'O': '6',
	'P': '7',
	'Q': '7',
	'R': '7',
	'S': '7',
	'T': '8',
	'U': '8',
	'V': '8',
	'W': '9',
	'X': '9',
	'Y': '9',
	'Z': '9',
}

func main() {
	in, _ := os.Open("10921.in")
	defer in.Close()
	out, _ := os.Create("10921.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		var phoneNumber strings.Builder
		for i := range line {
			if digit, ok := dict[line[i]]; ok {
				phoneNumber.WriteByte(digit)
			} else {
				phoneNumber.WriteByte(line[i])
			}
		}
		fmt.Fprintln(out, phoneNumber.String())
	}
}
