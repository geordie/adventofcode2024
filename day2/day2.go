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

type Game []Round

func Day2Puzzle1() {

	games := parseGames()

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

}

func Day2Puzzle2() {
	iAnswer := 0
	games := parseGames()
	for _, game := range games {
		idealRound := game.idealRound()
		iPower := idealRound.power()
		iAnswer += iPower
	}
	fmt.Println("Answer: ", iAnswer)
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

func parseGames() []Game {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	games := []Game{}
	iGameCounter := 0
	for scanner.Scan() {
		iGameCounter++
		s := scanner.Text()
		_, sRounds, bFound := strings.Cut(s, ": ")
		if !bFound {
			os.Exit(2)
		}

		arrRounds := strings.Split(sRounds, "; ")

		game := Game{}
		for _, sRound := range arrRounds {
			round := parseRound(sRound)
			game = append(game, round)
		}
		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return games
}

func (g Game) idealRound() Round {
	idealRound := Round{}
	for _, round := range g {
		idealRound.red = max(idealRound.red, round.red)
		idealRound.green = max(idealRound.green, round.green)
		idealRound.blue = max(idealRound.blue, round.blue)
	}
	return idealRound
}

func (r Round) power() int {
	return r.red * r.green * r.blue
}
