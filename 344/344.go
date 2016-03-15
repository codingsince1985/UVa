// UVa 344 - Roman Digititis

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	num = map[int]string{
		0:   "",
		1:   "i",
		2:   "ii",
		3:   "iii",
		4:   "iv",
		5:   "v",
		6:   "vi",
		7:   "vii",
		8:   "viii",
		9:   "ix",
		10:  "x",
		20:  "xx",
		30:  "xxx",
		40:  "xl",
		50:  "l",
		60:  "lx",
		70:  "lxx",
		80:  "lxxx",
		90:  "xc",
		100: "c",
	}
	romes = []byte{'i', 'v', 'x', 'l', 'c'}
)

func solve(n int) string {
	if str, ok := num[n]; ok {
		return str
	}
	return num[n/10*10] + num[n%10]
}

func output(out *os.File, n int) {
	dict := make(map[byte]int)
	for i := 1; i <= n; i++ {
		for _, v := range solve(i) {
			dict[byte(v)]++
		}
	}
	stats := make([]string, 5)
	for i, r := range romes {
		stats[i] = fmt.Sprintf("%d %c", dict[r], r)
	}
	fmt.Fprintf(out, "%d: %s\n", n, strings.Join(stats, ", "))
}

func main() {
	in, _ := os.Open("344.in")
	defer in.Close()
	out, _ := os.Create("344.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		output(out, n)
	}
}
