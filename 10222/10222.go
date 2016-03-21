// UVa 10222 - Decode the Mad man

package main

import (
	"fmt"
	"os"
)

func buildMap() map[byte]byte {
	keys := "`1234567890-=qwertyuiop[]\\asdfghjkl;'zxcvbnm,./"
	dict := make(map[byte]byte)
	for i := range keys {
		if i >= 2 {
			v := keys[i]
			dict[v] = keys[i-2]
			if v >= 'a' && v <= 'z' {
				dict['A'+v-'a'] = keys[i-2]
			}
		}
	}
	return dict
}

func main() {
	in, _ := os.Open("10222.in")
	defer in.Close()
	out, _ := os.Create("10222.out")
	defer out.Close()

	dict := buildMap()
	var ch byte
	for {
		if _, err := fmt.Fscanf(in, "%c", &ch); err != nil {
			break
		}
		if c, ok := dict[ch]; ok {
			ch = c
		}
		fmt.Fprintf(out, "%c", ch)
	}
}
