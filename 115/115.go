// UVa 115 - Climbing Trees

package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var a = make(map[string]map[string]bool)

func ancestor(n string, level int) map[string]int {
	as := make(map[string]int)
	for i := range a[n] {
		as[i] = level
		for j, vj := range ancestor(i, level+1) {
			as[j] = vj
		}
	}
	return as
}

func great(n int) string {
	if n > 0 {
		return strings.Repeat("great ", n)
	}
	return ""
}

func grand(a bool) string {
	if a {
		return "grand "
	}
	return ""
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findLow(a1, a2 map[string]int) int {
	l := math.MaxInt32
	for i := range a1 {
		if tmp, ok := a2[i]; ok {
			if tmp < l {
				l = tmp
			}
		}
	}
	return l
}

func check(n1, n2 string) string {
	a1, a2 := ancestor(n1, 0), ancestor(n2, 0)
	switch l1, l2 := findLow(a1, a2), findLow(a2, a1); {
	case l1 == math.MaxInt32:
		return "no relation"
	case l1 == 0 && l2 == 0:
		return "sibling"
	case l2 == 0:
		return great(l1-l2-2) + grand(l1 > l2+1) + "parent"
	case l1 == 0:
		return great(l2-l1-2) + grand(l2 > l1+1) + "child"
	default:
		return fmt.Sprintf("%d cousin removed %d", min(l1, l2), abs(l1-l2))
	}
}

func main() {
	in, _ := os.Open("115.in")
	defer in.Close()
	out, _ := os.Create("115.out")
	defer out.Close()

	var n1, n2 string
	for {
		if fmt.Fscanf(in, "%s%s", &n1, &n2); n1 == "no.child" {
			break
		}
		if _, ok := a[n1]; !ok {
			a[n1] = make(map[string]bool)
		}
		a[n1][n2] = true
	}
	for {
		if _, err := fmt.Fscanf(in, "%s%s", &n1, &n2); err != nil {
			break
		}
		fmt.Fprintln(out, check(n1, n2))
	}
}
