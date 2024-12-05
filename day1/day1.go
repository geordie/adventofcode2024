package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func Day1Puzzle1() {

	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	intsFirst := make([]int, 0)
	intsSecond := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		re := regexp.MustCompile("[0-9]+")
		sIntTokens := re.FindAllString(s, -1)
		sFirst := sIntTokens[0]
		sLast := sIntTokens[1]
		iFirst, err := strconv.Atoi(sFirst)
		if err != nil {
			log.Fatal(err)
		}
		iLast, err := strconv.Atoi(sLast)
		if err != nil {
			log.Fatal(err)
		}
		intsFirst = append(intsFirst, iFirst)
		intsSecond = append(intsSecond, iLast)
	}
	sort.Slice(intsFirst, func(i, j int) bool {
		return intsFirst[i] < intsFirst[j]
	})

	sort.Slice(intsSecond, func(i, j int) bool {
		return intsSecond[i] < intsSecond[j]
	})

	intsDiffs := make([]int, 0)
	for i := 0; i < len(intsFirst); i++ {
		iDiff := intsSecond[i] - intsFirst[i]
		if iDiff < 0 {
			iDiff = iDiff * -1
		}
		intsDiffs = append(intsDiffs, iDiff)
	}

	iAnswer := 0
	for _, i := range intsDiffs {
		iAnswer += i
	}

	fmt.Println(iAnswer)
}

func Day1Puzzle2() {

	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	intsFirst := make(map[int]int)
	intsSecond := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := scanner.Text()
		re := regexp.MustCompile("[0-9]+")
		sIntTokens := re.FindAllString(s, -1)
		sFirst := sIntTokens[0]
		sLast := sIntTokens[1]

		iFirst, err := strconv.Atoi(sFirst)
		if err != nil {
			log.Fatal(err)
		}

		iLast, err := strconv.Atoi(sLast)
		if err != nil {
			log.Fatal(err)
		}

		val, exists := intsFirst[iFirst]
		if !exists {
			intsFirst[iFirst] = 0
		} else {
			iOccurrences := val
			intsFirst[iFirst] = iOccurrences + 1
		}

		intsSecond = append(intsSecond, iLast)
	}

	sort.Slice(intsSecond, func(i, j int) bool {
		return intsSecond[i] < intsSecond[j]
	})

	for i := 0; i < len(intsSecond); i++ {
		iSecond := intsSecond[i]
		val, exists := intsFirst[iSecond]
		if exists {
			iOccurrences := val
			intsFirst[iSecond] = iOccurrences + 1
		}
	}

	iAnswer := 0
	for k, v := range intsFirst {
		if v > 0 {
			iAnswer += k * v
			fmt.Println(k, v)
		}
	}

	fmt.Println(iAnswer)
}
