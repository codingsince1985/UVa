// UVa 183 - Bit Maps

package main

import (
	"fmt"
	"os"
	"strings"
)

const max = 50

var in, out *os.File

func split(r, c int) (int, int, int, int) {
	r1, r2, c1, c2 := r/2, r/2, c/2, c/2
	if r%2 == 1 {
		r1++
	}
	if c%2 == 1 {
		c1++
	}
	return r1, r2, c1, c2
}

func same(bits [][]byte) bool {
	ch := bits[0][0]
	for _, row := range bits {
		if string(row) != strings.Repeat(string(ch), len(row)) {
			return false
		}
	}
	return true
}

func cut(y, x, h, l int, bits [][]byte) [][]byte {
	part := make([][]byte, h)
	for i := range part {
		part[i] = make([]byte, l)
		copy(part[i], bits[y+i][x:x+l])
	}
	return part
}

func solveB(bits [][]byte) string {
	if len(bits) > 0 && len(bits[0]) > 0 {
		if same(bits) {
			return string(bits[0][0])
		}
		r1, r2, c1, c2 := split(len(bits), len(bits[0]))
		return "D" + solveB(cut(0, 0, r1, c1, bits)) + solveB(cut(0, c1, r1, c2, bits)) +
			solveB(cut(r1, 0, r2, c1, bits)) + solveB(cut(r1, c1, r2, c2, bits))
	}
	return ""
}

func assemble(ul, ur, ll, lr [][]byte) [][]byte {
	bits := make([][]byte, len(ul)+len(ll))
	for i := range bits {
		bits[i] = make([]byte, len(ul[0]))
		if i < len(ul) {
			copy(bits[i], ul[i])
			bits[i] = append(bits[i], ur[i]...)
		} else {
			copy(bits[i], ll[i-len(ul)])
			bits[i] = append(bits[i], lr[i-len(ul)]...)
		}
	}
	return bits
}

func fill(r, c int, ch byte) [][]byte {
	bits := make([][]byte, r)
	for i := range bits {
		bits[i] = []byte(strings.Repeat(string(ch), c))
	}
	return bits
}

func solveD(r, c int) [][]byte {
	var bits [][]byte
	if r > 0 && c > 0 {
		var ch byte
		switch fmt.Fscanf(in, "%c", &ch); ch {
		case 'D':
			r1, r2, c1, c2 := split(r, c)
			ul, ur, ll, lr := solveD(r1, c1), solveD(r1, c2), solveD(r2, c1), solveD(r2, c2)
			bits = assemble(ul, ur, ll, lr)
		case '0', '1':
			bits = fill(r, c, ch)
		}
	}
	return bits
}

func output(bits [][]byte) {
	var line string
	for _, row := range bits {
		line += string(row)
	}
	for len(line) >= max {
		fmt.Fprintln(out, line[:max])
		line = line[max:]
	}
	if len(line) > 0 {
		fmt.Fprintln(out, line)
	}
}

func main() {
	in, _ = os.Open("183.in")
	defer in.Close()
	out, _ = os.Create("183.out")
	defer out.Close()

	var ty byte
	var r, c int
	for {
		if fmt.Fscanf(in, "%c", &ty); ty == '#' {
			break
		}
		fmt.Fscanf(in, "%d%d", &r, &c)
		switch ty {
		case 'B':
			fmt.Fprintf(out, "D%4d%4d\n", r, c)
			var m, line string
			for len(m) < r*c {
				fmt.Fscanf(in, "%s", &line)
				m += line
			}
			bits := make([][]byte, r)
			for i := range bits {
				bits[i] = []byte(m[i*c : (i+1)*c])
			}
			fmt.Fprintln(out, solveB(bits))
		case 'D':
			fmt.Fprintf(out, "B%4d%4d\n", r, c)
			output(solveD(r, c))
			fmt.Fscanf(in, "%c", &ty)
		}
	}
}
