package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Report struct {
	Levels []int
	Deltas []int
}

func Day2Puzzle1() {

	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	iAnswer := 0

	for scanner.Scan() {
		s := scanner.Text()
		report := Report{}
		report.Parse(s)
		if report.IsSafe(false) {
			iAnswer++
		}
	}

	fmt.Println(iAnswer)
}

func Day2Puzzle2() {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	iAnswer := 0

	for scanner.Scan() {
		s := scanner.Text()
		report := Report{}
		report.Parse(s)
		if report.IsSafe(true) {
			iAnswer++
		}
	}

	fmt.Println(iAnswer)
}

func (r *Report) Parse(sLevels string) {
	re := regexp.MustCompile("[0-9]+")
	sIntTokens := re.FindAllString(sLevels, -1)
	for _, s := range sIntTokens {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		r.Levels = append(r.Levels, i)
	}

	for i := 0; i < len(r.Levels)-1; i++ {
		iDiff := r.Levels[i+1] - r.Levels[i]
		r.Deltas = append(r.Deltas, iDiff)
	}
}

func (r *Report) IsSafe(bTest2 bool) bool {

	bSafe := r.LevelsDecreasing() || r.LevelsIncreasing()
	if !bSafe && bTest2 {
		bSafe = r.IsSafe2()
	}
	return bSafe
}

func (r *Report) IsSafe2() bool {

	bResult := false
	for i := 0; i < len(r.Levels); i++ {
		subReport := Report{}
		subReportLevels := make([]int, 0)
		for j := 0; j < len(r.Levels); j++ {
			if i == j {
				continue
			}
			subReportLevels = append(subReportLevels, r.Levels[j])
			subReport.Levels = subReportLevels
		}
		if subReport.IsSafe(false) {
			bResult = true
			break
		}
	}
	return bResult
}

func (r *Report) LevelsDecreasing() bool {

	if r.Levels == nil || len(r.Levels) < 1 {
		return false
	}

	bResult := true

	for i := 0; i < len(r.Levels)-1; i++ {
		iDiff := r.Levels[i+1] - r.Levels[i]
		if iDiff > -1 || iDiff < -3 {
			bResult = false
			break
		}
	}

	return bResult
}

func (r *Report) LevelsIncreasing() bool {

	if r.Levels == nil || len(r.Levels) < 1 {
		return false
	}

	bResult := true

	for i := 0; i < len(r.Levels)-1; i++ {
		iDiff := r.Levels[i+1] - r.Levels[i]
		if iDiff < 1 || iDiff > 3 {
			bResult = false
			break
		}
	}

	return bResult
}
