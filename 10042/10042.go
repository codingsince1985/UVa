// UVa 10042 - Smith Numbers

package main

import (
	"fmt"
	"os"
)

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return n == 2
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactorize(n int) []int {
	var p []int
	for i := 2; i <= n; i++ {
		for isPrime(i) {
			for n%i == 0 {
				p = append(p, i)
				n /= i
			}
			break
		}
	}
	return p
}

func main() {
	in, _ := os.Open("10042.in")
	defer in.Close()
	out, _ := os.Create("10042.out")
	defer out.Close()

	var n, num int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &num)
		for {
			sum1 := digitSum(num)
			sum2 := 0
			for _, v := range primeFactorize(num) {
				sum2 += digitSum(v)
			}
			if sum1 == sum2 {
				fmt.Fprintln(out, num)
				break
			}
			num++
		}
	}
}
