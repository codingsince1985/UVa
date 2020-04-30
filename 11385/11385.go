// UVa 11385 - Da Vinci Code

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var keepLetter = func(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r
	}
	return -1
}

var fibonacciMap = func() map[int]int {
	fib := []int{1, 2}
	for idx := 2; ; idx++ {
		if next := fib[idx-1] + fib[idx-2]; next < math.MaxInt32 {
			fib = append(fib, next)
		} else {
			break
		}
	}
	fibMap := make(map[int]int)
	for i, f := range fib {
		fibMap[f] = i + 1
	}
	return fibMap
}()

func solve(nums []int, line string) string {
	output := make([]byte, 100)
	for i := range output {
		output[i] = ' '
	}
	line = strings.Map(keepLetter, line)
	for i, num := range nums {
		idx := fibonacciMap[num]
		output[idx-1] = line[i]
	}
	return strings.TrimSpace(string(output))
}

func main() {
	in, _ := os.Open("11385.in")
	defer in.Close()
	out, _ := os.Create("11385.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	for t, _ := strconv.Atoi(s.Text()); t > 0 && s.Scan(); t-- {
		n, _ := strconv.Atoi(s.Text())
		nums := make([]int, n)
		s.Scan()
		r := strings.NewReader(s.Text())
		for i := range nums {
			fmt.Fscanf(r, "%d", &nums[i])
		}
		s.Scan()
		fmt.Fprintln(out, solve(nums, s.Text()))
	}
}
