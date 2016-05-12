// UVa 536 - Tree Recovery

package main

import (
	"fmt"
	"os"
)

func index(n byte, lst []byte) int {
	for i := range lst {
		if n == lst[i] {
			return i
		}
	}
	return -1
}

func post(preord, inord []byte) []byte {
	if len(preord) == 0 {
		return []byte{}
	}
	idx := index(preord[0], inord)
	return append(append(post(preord[1:idx+1], inord[0:idx]), post(preord[idx+1:], inord[idx+1:])...), preord[0])
}

func main() {
	in, _ := os.Open("536.in")
	defer in.Close()
	out, _ := os.Create("536.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		preord := []byte(line)
		fmt.Fscanf(in, "%s", &line)
		inord := []byte(line)
		postord := post(preord, inord)
		for i := range postord {
			fmt.Fprintf(out, "%c", postord[i])
		}
		fmt.Fprintln(out)
	}
}
