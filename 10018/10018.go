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

	var n int
	var p uint32
	fmt.Fscanf(in, "%d", &n)
	for i := 0; i < n; i++ {
		count := 0
		fmt.Fscanf(in, "%d", &p)
		for {
			if isPalindrome(p) {
				break
			}
			p += reverse(p)
			count++
		}
		fmt.Fprintln(out, count, p)
	}
}
