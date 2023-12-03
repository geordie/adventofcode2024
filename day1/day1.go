package calories

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/geordie/adventofcode2023/util"
)

func Day1Puzzle1() {

	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	ints := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		re := regexp.MustCompile("[1-9]")
		sIntTokens := re.FindAllString(s, -1)
		sFirst := sIntTokens[0]
		iFirst, _ := strconv.Atoi(sFirst)
		sLast := sIntTokens[len(sIntTokens)-1]
		iLast, _ := strconv.Atoi(sLast)

		coord := (iFirst * 10) + iLast

		fmt.Println(s, sIntTokens, coord)

		ints = append(ints, coord)
	}

	fmt.Println(ints)
	sum := int32(0)
	for _, num := range ints {
		sum += int32(num)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Day1Puzzle2() {

	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	iAnswer := 0

	scanner := bufio.NewScanner(file)
	numTokens := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		s := scanner.Text()

		iMinFirstIndex := len(s)
		iFirst := -1

		for _, item := range numTokens {
			idx := strings.Index(s, item)
			if idx > -1 && idx < iMinFirstIndex {
				iMinFirstIndex = idx
				if len(item) == 1 {
					iFirst, err = strconv.Atoi(item)
				} else {
					iFirst = util.GetIntFromEnglish(item)
				}
			}
		}

		iMaxLastIndex := -1
		iLast := -1
		for _, item := range numTokens {
			idx := strings.LastIndex(s, item)
			if idx > -1 && idx > iMaxLastIndex {
				iMaxLastIndex = idx
				if len(item) == 1 {
					iLast, err = strconv.Atoi(item)
				} else {
					iLast = util.GetIntFromEnglish(item)
				}
			}
		}

		coord := (iFirst * 10) + iLast
		iAnswer += coord
	}

	fmt.Println(iAnswer)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
