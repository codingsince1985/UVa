// UVa 10018 - Reverse and Add

package main

import (
	"fmt"
	"os"
)

func isPalindrome(num uint32) bool {
	str := fmt.Sprint(num)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func reverse(num uint32) uint32 {
	var result uint32
	for num != 0 {
		result *= 10
		result += num % 10
		num /= 10
	}
	return result
}

func main() {
	in, _ := os.Open("10018.in")
	defer in.Close()
	out, _ := os.Create("10018.out")
	defer out.Close()

	var n, count int
	var p uint32
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &p)
		for count = 0; ; count++ {
			if isPalindrome(p) {
				break
			}
			p += reverse(p)
		}
		fmt.Fprintln(out, count, p)
	}
}
