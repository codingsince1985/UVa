// UVa 10101 - Bangla Numbers

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	kuti  = 10000000
	lakh  = 100000
	hajar = 1000
	shata = 100
)

func split(line string) []int {
	var num int
	var nums []int
	for len(line) > 0 {
		if len(line) >= 7 {
			fmt.Sscanf(line[len(line)-7:], "%d", &num)
			line = line[:len(line)-7]
		} else {
			fmt.Sscanf(line, "%d", &num)
			line = ""
		}
		nums = append([]int{num}, nums...)
	}
	return nums
}

func convert(n int) []string {
	var tokens []string
	for n > 0 {
		switch l := len(strconv.Itoa(n)); {
		case l > 7:
			tokens = append(tokens, strconv.Itoa(n/kuti)+" kuti")
			n %= kuti
		case l > 5:
			tokens = append(tokens, strconv.Itoa(n/lakh)+" lakh")
			n %= lakh
		case l > 3:
			tokens = append(tokens, strconv.Itoa(n/hajar)+" hajar")
			n %= hajar
		case l > 2:
			tokens = append(tokens, strconv.Itoa(n/shata)+" shata")
			n %= shata
		default:
			tokens = append(tokens, strconv.Itoa(n))
			n = 0
		}
	}
	return tokens
}

func main() {
	in, _ := os.Open("10101.in")
	defer in.Close()
	out, _ := os.Create("10101.out")
	defer out.Close()

	var num string
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%s", &num); err != nil {
			break
		}
		if num == "0" {
			fmt.Fprintf(out, "% 4d: 0\n", kase)
		} else {
			nums := split(num)
			var tokens []string
			for i, vi := range nums {
				if i > 0 {
					tokens[len(tokens)-1] = tokens[len(tokens)-1] + " kuti"
				}
				tokens = append(tokens, convert(vi)...)
			}
			fmt.Fprintf(out, "% 4d: %s\n", kase, strings.Join(tokens, " "))
		}
	}
}
