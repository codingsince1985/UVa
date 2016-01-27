// UVa 10194 - Football (aka Soccer)

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
	"sort"
)

var s *bufio.Scanner
var teamMap map[string]team

type team struct {
	name                              string
	points                            int
	games, wins, ties, losses         int
	goalDiff, goalScored, goalAgainst int
}

type teams []team

func cmp(t1, t2 team) bool {
	if t1.points != t2.points {
		return t1.points > t2.points
	} else if t1.wins != t2.wins {
		return t1.wins > t2.wins
	} else if t1.goalDiff != t2.goalDiff {
		return t1.goalDiff > t2.goalDiff
	} else if t1.goalScored != t2.goalScored {
		return t1.goalScored > t2.goalScored
	} else if t1.games != t2.games {
		return t1.games < t2.games
	} else {
		return strings.Compare(t1.name, t2.name) == -1
	}
}

func (t teams)Len() int {
	return len(t)
}

func (t teams)Less(i, j int) bool {
	return cmp(t[i], t[j])
}

func (t teams)Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func nextLine() string {
	s.Scan()
	return s.Text()
}

func values(m map[string]team) teams {
	var v []team
	for _, value := range m {
		v = append(v, value)
	}
	return v
}

func process(a, b string, s1, s2 int) {
	t1, t2 := teamMap[a], teamMap[b]
	t1.games++
	t2.games++
	t1.goalScored += s1
	t1.goalAgainst += s2
	t2.goalScored += s2
	t2.goalAgainst += s1
	t1.goalDiff = t1.goalScored - t1.goalAgainst
	t2.goalDiff = t2.goalScored - t2.goalAgainst
	if s1 > s2 {
		t1.wins ++
		t1.points += 3
		t2.losses ++
	} else if s2 > s1 {
		t1.losses ++
		t2.wins ++
		t2.points += 3
	} else {
		t1.ties ++
		t1.points ++
		t2.ties ++
		t2.points ++
	}
	teamMap[a], teamMap[b] = t1, t2
}

func output(out *os.File, t teams) {
	for i, v := range t {
		fmt.Fprintf(out, "%d) %s %dp, %dg (%d-%d-%d), %dgd (%d-%d)\n",
			i + 1, v.name, v.points, v.games, v.wins, v.ties, v.losses, v.goalDiff, v.goalScored, v.goalAgainst)
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("10194.in")
	defer in.Close()
	out, _ := os.Create("10194.out")
	defer out.Close()

	s = bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	n, _ := strconv.Atoi(nextLine())
	for i := 0; i < n; i++ {
		teamMap = make(map[string]team)
		fmt.Fprintln(out, nextLine())

		tm, _ := strconv.Atoi(nextLine())
		for j := 0; j < tm; j++ {
			tmp := nextLine()
			teamMap[tmp] = team{name:tmp}
		}

		gm, _ := strconv.Atoi(nextLine())
		for j := 0; j < gm; j++ {
			tokens := strings.Split(nextLine(), "#")
			scores := strings.Split(tokens[1], "@")
			s1, _ := strconv.Atoi(scores[0])
			s2, _ := strconv.Atoi(scores[1])
			process(tokens[0], tokens[2], s1, s2)
		}

		v := values(teamMap)
		sort.Sort(v)
		output(out, v)
	}
}