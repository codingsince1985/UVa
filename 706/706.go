// UVa 706 - LC-Display

package main

import (
	"fmt"
	"os"
)

var digits = [][]byte{
	{'-', '|', '|', ' ', '|', '|', '-'},
	{' ', ' ', '|', ' ', ' ', '|', ' '},
	{'-', ' ', '|', '-', '|', ' ', '-'},
	{'-', ' ', '|', '-', ' ', '|', '-'},
	{' ', '|', '|', '-', ' ', '|', ' '},
	{'-', '|', ' ', '-', ' ', '|', '-'},
	{'-', '|', ' ', '-', '|', '|', '-'},
	{'-', ' ', '|', ' ', ' ', '|', ' '},
	{'-', '|', '|', '-', '|', '|', '-'},
	{'-', '|', '|', '-', ' ', '|', '-'},
}

func output135(out *os.File, n int, num string, idx int) {
	for i := range num {
		d := num[i] - '0'
		fmt.Fprint(out, " ")
		for j := 0; j < n; j++ {
			fmt.Fprintf(out, "%c", digits[d][idx])
		}
		fmt.Fprint(out, " ")
		if i != len(num)-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}

func output24(out *os.File, n int, num string, idx int) {
	for i := 0; i < n; i++ {
		for j := range num {
			d := num[j] - '0'
			fmt.Fprintf(out, "%c", digits[d][idx])
			for k := 0; k < n; k++ {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, "%c", digits[d][idx+1])
			if j != len(num)-1 {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}
}

func output(out *os.File, n int, num string) {
	output135(out, n, num, 0)
	output24(out, n, num, 1)
	output135(out, n, num, 3)
	output24(out, n, num, 4)
	output135(out, n, num, 6)
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("706.in")
	defer in.Close()
	out, _ := os.Create("706.out")
	defer out.Close()

	var n int
	var num string
	for {
		if fmt.Fscanf(in, "%d%s", &n, &num); n == 0 {
			break
		}
		output(out, n, num)
	}
}
