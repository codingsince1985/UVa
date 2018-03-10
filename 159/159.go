// UVa 159 - Word Crosses

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type cross struct{ h, v int }

var noCross = cross{-1, -1}

func leadingWordCross(h, v string) cross {
	for i := range h {
		if idx := strings.Index(v, string(h[i])); idx >= 0 {
			return cross{i, idx}
		}
	}
	return cross{-1, -1}
}

func solve(out io.Writer, h1, v1, h2, v2 string) {
	c1, c2 := leadingWordCross(h1, v1), leadingWordCross(h2, v2)
	if c1 == noCross || c2 == noCross {
		fmt.Fprintln(out, "Unable to make two crosses")
		return
	}
	var idx1, idx2 int
	switch {
	case c1.v > c2.v:
		for i := 0; i < (c1.v - c2.v); i++ {
			fmt.Fprintf(out, "%s%c\n", strings.Repeat(" ", c1.h), v1[idx1])
			idx1++
		}
	case c1.v < c2.v:
		for i := 0; i < (c2.v - c1.v); i++ {
			fmt.Fprintf(out, "%s%c\n", strings.Repeat(" ", len(h1)+3+c2.h), v2[idx2])
			idx2++
		}
	}
	for idx1 < c1.v {
		fmt.Fprintf(out, "%s%c%s%c\n", strings.Repeat(" ", c1.h), v1[idx1], strings.Repeat(" ", len(h1)-c1.h-1+3+c2.h), v2[idx2])
		idx1++
		idx2++
	}
	fmt.Fprintf(out, "%s   %s\n", h1, h2)
	idx1++
	idx2++
	for idx1 < len(v1) || idx2 < len(v2) {
		switch {
		case idx1 < len(v1) && idx2 < len(v2):
			fmt.Fprintf(out, "%s%c%s%c\n", strings.Repeat(" ", c1.h), v1[idx1], strings.Repeat(" ", len(h1)-c1.h-1+3+c2.h), v2[idx2])
			idx1++
			idx2++
		case idx1 < len(v1):
			fmt.Fprintf(out, "%s%c\n", strings.Repeat(" ", c1.h), v1[idx1])
			idx1++
		case idx2 < len(v2):
			fmt.Fprintf(out, "%s%c\n", strings.Repeat(" ", len(h1)+3+c2.h), v2[idx2])
			idx2++
		}
	}
}

func main() {
	in, _ := os.Open("159.in")
	defer in.Close()
	out, _ := os.Create("159.out")
	defer out.Close()

	var h1, v1, h2, v2 string
	for first := true; ; {
		if fmt.Fscanf(in, "%s", &h1); h1 == "#" {
			break
		}
		fmt.Fscanf(in, "%s%s%s", &v1, &h2, &v2)
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(out, h1, v1, h2, v2)
	}
}
