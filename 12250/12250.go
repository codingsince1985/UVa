// UVa 12250 - Language Detection

package main

import (
	"fmt"
	"os"
)

var wordMap = map[string]string{"HELLO": "ENGLISH", "HOLA": "SPANISH", "HALLO": "GERMAN", "BONJOUR": "FRENCH", "CIAO": "ITALIAN", "ZDRAVSTVUJTE": "RUSSIAN"}

func solve(word string) string {
	if country, ok := wordMap[word]; ok {
		return country
	}
	return "UNKNOWN"
}

func main() {
	in, _ := os.Open("12250.in")
	defer in.Close()
	out, _ := os.Create("12250.out")
	defer out.Close()

	var word string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%s", &word); word == "#" {
			break
		}
		fmt.Fprintf(out, "Case %d: %s\n", kase, solve(word))
	}
}
