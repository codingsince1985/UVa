// UVa 10700 - Camel trading

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func multiply(toMul []string) int64 {
	var product int64 = 1
	for _, v := range toMul {
		token, _ := strconv.ParseInt(v, 10, 64)
		product *= token
	}
	return product
}

func findMin(line string) int64 {
	toAdd := strings.Split(line, "+")
	for i, vi := range toAdd {
		toMul := strings.Split(vi, "*")
		toAdd[i] = fmt.Sprintf("%d", multiply(toMul))
	}
	return add(toAdd)
}

func add(toAdd []string) int64 {
	var total int64
	for _, v := range toAdd {
		token, _ := strconv.ParseInt(v, 10, 64)
		total += token
	}
	return total
}

func findMax(line string) int64 {
	toMul := strings.Split(line, "*")
	for i, vi := range toMul {
		toAdd := strings.Split(vi, "+")
		toMul[i] = fmt.Sprintf("%d", add(toAdd))
	}
	return multiply(toMul)
}

func main() {
	in, _ := os.Open("10700.in")
	defer in.Close()
	out, _ := os.Create("10700.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	for s.Scan() && n > 0 {
		line := s.Text()
		fmt.Fprintf(out, "The maximum and minimum are %d and %d.\n", findMax(line), findMin(line))
		n--
	}
}
