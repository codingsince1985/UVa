// UVa 10924 - Prime Words

package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeWord(word string) bool {
	total := 0
	for i := range word {
		if word[i] >= 'a' && word[i] <= 'z' {
			total += int(word[i] - 'a' + 1)
		} else {
			total += int(word[i] - 'A' + 27)
		}
	}
	return isPrime(total)
}

func main() {
	in, _ := os.Open("10924.in")
	defer in.Close()
	out, _ := os.Create("10924.out")
	defer out.Close()

	var word string
	for {
		if _, err := fmt.Fscanf(in, "%s", &word); err != nil {
			break
		}
		if isPrimeWord(word) {
			fmt.Fprintln(out, "It is a prime word")
		} else {
			fmt.Fprintln(out, "It is not a prime word.")
		}
	}
}
