package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	red   int
	green int
	blue  int
}

func Day2Puzzle1() {

	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	games := [][]Round{}
	iGameCounter := 0
	for scanner.Scan() {
		iGameCounter++
		s := scanner.Text()
		_, sRounds, bFound := strings.Cut(s, ": ")
		if !bFound {
			os.Exit(2)
		}

		arrRounds := strings.Split(sRounds, "; ")

		rounds := []Round{}
		for _, sRound := range arrRounds {
			round := parseRound(sRound)
			rounds = append(rounds, round)
		}
		games = append(games, rounds)
		fmt.Println(rounds)
	}

	iMaxRed := 12
	iMaxGreen := 13
	iMaxBlue := 14
	iAnswer := 0

	for idx, game := range games {
		iAnswer += (idx + 1)
		for _, round := range game {
			if round.red > iMaxRed ||
				round.green > iMaxGreen ||
				round.blue > iMaxBlue {
				iAnswer -= (idx + 1)
				break
			}
		}
	}

	fmt.Println(iAnswer)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseRound(sRound string) Round {
	arrSets := strings.Split(sRound, ", ")

	round := Round{}

	for _, colourCount := range arrSets {
		count, colour, _ := strings.Cut(colourCount, " ")
		if colour == "red" {
			round.red, _ = strconv.Atoi(count)
		} else if colour == "green" {
			round.green, _ = strconv.Atoi(count)
		} else if colour == "blue" {
			round.blue, _ = strconv.Atoi(count)
		}
	}
	return round
}
