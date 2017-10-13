// UVa 11223 - O: dah dah dah!

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var morse = map[string]byte{
	".-":     'A',
	"-...":   'B',
	"-.-.":   'C',
	"-..":    'D',
	".":      'E',
	"..-.":   'F',
	"--.":    'G',
	"....":   'H',
	"..":     'I',
	".---":   'J',
	"-.-":    'K',
	".-..":   'L',
	"--":     'M',
	"-.":     'N',
	"---":    'O',
	".--.":   'P',
	"--.-":   'Q',
	".-.":    'R',
	"...":    'S',
	"-":      'T',
	"..-":    'U',
	"...-":   'V',
	".--":    'W',
	"-..-":   'X',
	"-.--":   'Y',
	"--..":   'Z',
	"-----":  '0',
	".----":  '1',
	"..---":  '2',
	"...--":  '3',
	"....-":  '4',
	".....":  '5',
	"-....":  '6',
	"--...":  '7',
	"---..":  '8',
	"----.":  '9',
	".-.-.-": '.',
	"--..--": ',',
	"..--..": '?',
	".----.": '\'',
	"-.-.--": '!',
	"-..-.":  '/',
	"-.--.":  '(',
	"-.--.-": ')',
	".-...":  '&',
	"---...": ':',
	"-.-.-.": ';',
	"-...-":  '=',
	".-.-.":  '+',
	"-....-": '-',
	"..--.-": '_',
	".-..-.": '"',
	".--.-.": '@',
}

func solve(line string) string {
	words := strings.Split(line, "  ")
	sentence := make([]string, len(words))
	for i, word := range words {
		letters := strings.Split(word, " ")
		chars := make([]byte, len(letters))
		for j, letter := range letters {
			chars[j] = morse[letter]
		}
		sentence[i] = string(chars)
	}
	return strings.Join(sentence, " ")
}

func main() {
	in, _ := os.Open("11223.in")
	defer in.Close()
	out, _ := os.Create("11223.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	for kase := 1; kase <= n && s.Scan(); kase++ {
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Message #%d\n%s\n", kase, solve(s.Text()))
	}
}
