// UVa 377 - Cowculations

package main

import (
	"fmt"
	"os"
)

var (
	charMap = map[byte]int{'V': 0, 'U': 1, 'C': 2, 'D': 3}
	numMap  = map[int]string{0: "V", 1: "U", 2: "C", 3: "D"}
)

func convert(n string) [8]int {
	var num [8]int
	for idx, i := 7, len(n)-1; i >= 0; idx, i = idx-1, i-1 {
		num[idx] = charMap[n[i]]
	}
	return num
}

func add(num1, num2 [8]int) [8]int {
	for carry, i := false, 7; i >= 0; i-- {
		num2[i] += num1[i]
		if carry {
			num2[i]++
			carry = false
		}
		if num2[i] > 3 {
			carry = true
			num2[i] -= 4
		}
	}
	return num2
}

func toString(num [8]int) string {
	var str string
	for i := range num {
		str += numMap[num[i]]
	}
	return str
}

func shiftLeft(num [8]int) [8]int {
	for i := 0; i < 7; i++ {
		num[i] = num[i+1]
	}
	num[7] = 0
	return num
}

func shiftRight(num [8]int) [8]int {
	for i := 7; i > 0; i-- {
		num[i] = num[i-1]
	}
	num[0] = 0
	return num
}

func calc(n1, n2, op string) string {
	num2 := convert(n2)
	switch op {
	case "A":
		num2 = add(convert(n1), num2)
	case "R":
		num2 = shiftRight(num2)
	case "L":
		num2 = shiftLeft(num2)
	}
	return toString(num2)
}

func main() {
	in, _ := os.Open("377.in")
	defer in.Close()
	out, _ := os.Create("377.out")
	defer out.Close()

	var kase int
	var op, n1, n2, result string
	fmt.Fprintln(out, "COWCULATIONS OUTPUT")
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &n1)
		fmt.Fscanf(in, "%s", &n2)
		for i := 0; i < 3; i++ {
			fmt.Fscanf(in, "%s", &op)
			n2 = calc(n1, n2, op)
		}
		if fmt.Fscanf(in, "%s", &result); n2 == result {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
	fmt.Fprintln(out, "END OF OUTPUT")
}
