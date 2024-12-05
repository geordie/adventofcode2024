package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day2Puzzle1() {

	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	iAnswer := 0
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	for scanner.Scan() {
		s := scanner.Text()
		sTokens := re.FindAllString(s, -1)
		for _, sToken := range sTokens {
			iProduct := FindProduct(sToken)
			iAnswer += iProduct
		}
	}

	fmt.Println(iAnswer)
}

type Deadzone struct {
	Start int
	End   int
}

func Day2Puzzle2() {

	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	iAnswer := 0
	reDos := regexp.MustCompile(`do\(\)`)
	reDonts := regexp.MustCompile(`don\'t\(\)`)
	reMuls := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	for scanner.Scan() {
		s := scanner.Text()

		idxDos := reDos.FindAllStringIndex(s, -1)
		idxDonts := reDonts.FindAllStringIndex(s, -1)

		fmt.Println("")
		fmt.Println("DOs: ", idxDos)
		fmt.Println("DONTs: ", idxDonts)
		doPtr := 0
		deadZones := make([]Deadzone, 0)
		for _, idxDont := range idxDonts {
			if idxDont[0] > doPtr {
				deadZone := Deadzone{Start: idxDont[0], End: len(s)}
				for _, idxDo := range idxDos {
					doPtr = idxDo[1]
					if idxDo[0] > idxDont[0] {
						deadZone.End = idxDo[0]
						break
					}
				}
				deadZones = append(deadZones, deadZone)
			}
		}
		fmt.Println(deadZones)

		idxMuls := reMuls.FindAllStringIndex(s, -1)
		for _, idxMul := range idxMuls {
			bDead := false
			for _, deadZone := range deadZones {
				if idxMul[0] > deadZone.Start && idxMul[0] < deadZone.End {
					bDead = true
					break
				}
			}
			if !bDead {
				sToken := s[idxMul[0]:idxMul[1]]
				iProduct := FindProduct(sToken)
				iAnswer += iProduct
			}
		}
	}

	fmt.Println(iAnswer)
}

// DAY 1: 161289189
// DAY 2: 89198456 IS TOO HIGH
// DAY 2: 24237394 IS TOO LOW
// DAY 2: 28208682 IS TOO LOW
// DAY 2: 85698778 IS WRONG
// DAY 2: 83595109 IS THE ANSWER, I SOLVED BY REGEXING THE FILE in TEXT EDITOR

func FindProduct(s string) int {
	iProduct := 1
	re := regexp.MustCompile("[0-9]+")
	sIntTokens := re.FindAllString(s, -1)
	for _, s := range sIntTokens {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		iProduct *= i
	}
	return iProduct
}
