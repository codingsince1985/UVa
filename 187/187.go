// UVa 187 - Transaction Processing

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	end          = "000"
	outOfBalance = "999"
)

var (
	out        *os.File
	accountMap = make(map[string]string)
)

type transaction struct {
	accountNumber string
	amount        int
}

func solve(sequence string, transactions []transaction) {
	var balance int
	for _, txn := range transactions {
		balance += txn.amount
	}
	if balance != 0 {
		accountMap[outOfBalance] = "Out of Balance"
		transactions = append(transactions, transaction{outOfBalance, -balance})
		fmt.Fprintf(out, "*** Transaction %s is out of balance ***\n", sequence)
		for _, txn := range transactions {
			fmt.Fprintf(out, "%s %-30s %10.2f\n", txn.accountNumber, accountMap[txn.accountNumber], float64(txn.amount)/100)
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in, _ := os.Open("187.in")
	defer in.Close()
	out, _ = os.Create("187.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	accountMap = make(map[string]string)
	for s.Scan() {
		if line = s.Text(); strings.HasPrefix(line, end) {
			break
		}
		accountMap[line[:3]] = line[3:]
	}
	var sequence string
	var transactions []transaction
	var amount int
	for s.Scan() {
		if line = s.Text(); strings.HasPrefix(line, end) {
			solve(sequence, transactions)
			break
		}
		if sequence != "" && sequence != line[:3] {
			solve(sequence, transactions)
			transactions = nil
		}
		sequence = line[:3]
		fmt.Sscanf(line[6:], "%d", &amount)
		transactions = append(transactions, transaction{line[3:6], amount})
	}
}
