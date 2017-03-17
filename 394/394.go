// UVa 394 - Mapmaker

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type (
	dimension struct{ lower, upper int }
	array     struct {
		base, bytes, dim int
		dimensions       []dimension
	}
)

func solve(line string, arrays map[string]array) string {
	var name string
	r := strings.NewReader(line)
	fmt.Fscanf(r, "%s", &name)
	arr := arrays[name]

	dims := make([]int, arr.dim)
	for i := range dims {
		fmt.Fscanf(r, "%d", &dims[i])
	}
	base, count := 1, 0
	for i := len(dims) - 1; i >= 0; i-- {
		count += base * (dims[i] - arr.dimensions[i].lower)
		base *= arr.dimensions[i].upper - arr.dimensions[i].lower + 1
	}
	return fmt.Sprintf("%s%s = %d", name, strings.Replace(fmt.Sprint(dims), " ", ", ", -1), arr.base+count*arr.bytes)
}

func main() {
	in, _ := os.Open("394.in")
	defer in.Close()
	out, _ := os.Create("394.out")
	defer out.Close()

	var n, r int
	var name string
	fmt.Fscanf(in, "%d%d", &n, &r)

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	arrays := make(map[string]array)
	for ; n > 0 && s.Scan(); n-- {
		var arr array
		r := strings.NewReader(s.Text())
		fmt.Fscanf(r, "%s%d%d%d", &name, &arr.base, &arr.bytes, &arr.dim)
		arr.dimensions = make([]dimension, arr.dim)
		for i := range arr.dimensions {
			fmt.Fscanf(r, "%d%d", &arr.dimensions[i].lower, &arr.dimensions[i].upper)
		}
		arrays[name] = arr
	}
	for ; r > 0 && s.Scan(); r-- {
		fmt.Fprintln(out, solve(s.Text(), arrays))
	}
}
