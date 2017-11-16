// UVa 486 - English-Number Translator

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	addMap = map[string]int{
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10,
		"eleven": 11, "twelve": 12, "thirteen": 13, "fourteen": 14, "fifteen": 15, "sixteen": 16, "seventeen": 17, "eighteen": 18, "nineteen": 19,
		"twenty": 20, "thirty": 30, "forty": 40, "fifty": 50, "sixty": 60, "seventy": 70, "eighty": 80, "ninety": 90,
	}
)

func calculate(line string) int {
	tokens := strings.Fields(strings.TrimSpace(line))
	var number int
	for _, token := range tokens {
		if token == "hundred" {
			number *= 100
		} else {
			number += addMap[token]
		}
	}
	return number
}

func calculate1K(line string) int {
	tokens := strings.Split(strings.TrimSpace(line), "thousand")
	var number int
	for i := range tokens {
		number = number*1000 + calculate(tokens[i])
	}
	return number
}

func calculate1M(line string) int {
	tokens := strings.Split(line, "million")
	var number int
	for _, token := range tokens {
		number = number*1000000 + calculate1K(token)
	}
	return number
}

func main() {
	in, _ := os.Open("486.in")
	defer in.Close()
	out, _ := os.Create("486.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		number := calculate1M(strings.Replace(line, "negative", " ", 1))
		if strings.HasPrefix(line, "negative") {
			number = -number
		}
		fmt.Fprintln(out, number)
	}
}
