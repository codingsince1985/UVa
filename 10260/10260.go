// UVa 10260 - Soundex

package main

import (
	"fmt"
	"os"
)

func buildDict() map[byte]int {
	chars := [][]byte{
		{},
		{'B', 'F', 'P', 'V'},
		{'C', 'G', 'J', 'K', 'Q', 'S', 'X', 'Z'},
		{'D', 'T'},
		{'L'},
		{'M', 'N'},
		{'R'},
	}

	dict := make(map[byte]int)
	for i, v := range chars {
		for _, b := range v {
			dict[b] = i
		}
	}
	return dict
}

func main() {
	in, _ := os.Open("10260.in")
	defer in.Close()
	out, _ := os.Create("10260.out")
	defer out.Close()

	dict := buildDict()
	var word string
	for {
		if _, err := fmt.Fscanf(in, "%s", &word); err != nil {
			break
		}
		var lastValue int
		for i := range word {
			if value := dict[word[i]]; value != lastValue {
				if value != 0 {
					fmt.Fprint(out, value)
				}
				lastValue = value
			}
		}
		fmt.Fprintln(out)
	}
}
