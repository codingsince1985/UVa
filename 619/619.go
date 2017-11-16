// UVa 619 - Numerically Speaking

package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

var (
	base = big.NewInt(26)
	zero = big.NewInt(0)
)

func split(number string) string {
	var nums []string
	if remain := len(number) % 3; remain != 0 {
		nums = append(nums, number[:remain])
		number = number[remain:]
	}
	for i := 0; i < len(number)/3; i++ {
		nums = append(nums, number[i*3:(i+1)*3])
	}
	return strings.Join(nums, ",")
}

func toNumber(line string) string {
	var num big.Int
	for i := range line {
		num.Mul(&num, base).Add(&num, big.NewInt(int64(line[i]-'a'+1)))
	}
	return split(num.String())
}

func toString(line string) string {
	var word string
	var num, mod big.Int
	num.SetString(line, 10)
	for num.Cmp(zero) > 0 {
		if num.DivMod(&num, base, &mod); mod.Int64() == 0 {
			word = "z" + word
		} else {
			word = string('a'+mod.Int64()-1) + word
		}
	}
	return word
}

func solve(line string) (string, string) {
	if line[0] >= '0' && line[0] <= '9' {
		return toString(line), split(line)
	}
	return line, toNumber(line)
}

func main() {
	in, _ := os.Open("619.in")
	defer in.Close()
	out, _ := os.Create("619.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%s", &line); line == "*" {
			break
		}
		word, num := solve(line)
		fmt.Fprintf(out, "%-20s  %s\n", word, num)
	}
}
