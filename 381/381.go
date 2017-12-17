// UVa 381 - Making the Grade

package main

import (
	"fmt"
	"math"
	"os"
)

type student struct {
	scores                []int
	bonus, absence, grade int
	average               float64
}

func round(n float64) float64 { return float64(int(n*10+.5)) / 10 }

func doMean(s, t int, students []student) float64 {
	var total float64
	for i, v := range students {
		sum, lowest := 0, 101
		for _, grade := range v.scores {
			sum += grade
			if grade < lowest {
				lowest = grade
			}
		}
		if t > 2 {
			students[i].average = round(float64(sum-lowest) / float64(t-1))
		} else {
			students[i].average = round(float64(sum) / float64(t))
		}
		total += students[i].average
	}
	return round(total / float64(s))
}

func doStandardDeviation(s int, mean float64, students []student) float64 {
	var sum float64
	for _, v := range students {
		sum += (v.average - mean) * (v.average - mean)
	}
	sum /= float64(s)
	if sd := round(math.Sqrt(sum)); sd >= 1 {
		return sd
	}
	return 1
}

func doBonus(students []student) {
	for i, v := range students {
		students[i].average += float64(int(v.bonus/2) * 3)
	}
}

func doGrade(mean, sd float64, students []student) {
	for i, v := range students {
		switch diff := v.average - mean; {
		case diff >= sd:
			students[i].grade = 4
		case diff >= 0:
			students[i].grade = 3
		case diff >= -sd:
			students[i].grade = 2
		default:
			students[i].grade = 1
		}
	}
}

func doAbsence(students []student) {
	for i, v := range students {
		if v.absence == 0 {
			if v.grade < 4 {
				students[i].grade++
			}
		} else {
			if students[i].grade -= v.absence / 4; students[i].grade < 0 {
				students[i].grade = 0
			}
		}
	}
}

func doClass(s int, students []student) float64 {
	var sum int
	for _, v := range students {
		sum += v.grade
	}
	return round(float64(sum) / float64(s))
}

func solve(s, t int, students []student) float64 {
	mean := doMean(s, t, students)
	sd := doStandardDeviation(s, mean, students)
	doBonus(students)
	doGrade(mean, sd, students)
	doAbsence(students)
	return doClass(s, students)
}

func main() {
	in, _ := os.Open("381.in")
	defer in.Close()
	out, _ := os.Create("381.out")
	defer out.Close()

	fmt.Fprintln(out, "MAKING THE GRADE OUTPUT")
	var n, s, t int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &s, &t)
		students := make([]student, s)
		for i := range students {
			students[i].scores = make([]int, t)
			for j := range students[i].scores {
				fmt.Fscanf(in, "%d", &students[i].scores[j])
			}
			fmt.Fscanf(in, "%d%d", &students[i].bonus, &students[i].absence)
		}
		fmt.Fprintf(out, "%.1f\n", solve(s, t, students))
	}
	fmt.Fprintln(out, "END OF OUTPUT")
}
