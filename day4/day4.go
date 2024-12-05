package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type XCoord struct {
	x int
	y int
}

type Board struct {
	lines []string
}

func Day4Puzzle1() {

	file, err := os.Open("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	reDos := regexp.MustCompile(`X`)
	scanner := bufio.NewScanner(file)
	iAnswer := 0
	iLine := 0
	xCoords := make([]XCoord, 0)
	lines := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		Xs := reDos.FindAllStringIndex(s, -1)
		for _, X := range Xs {
			xCoord := XCoord{x: X[0], y: iLine}
			xCoords = append(xCoords, xCoord)
		}
		lines = append(lines, s)
		iLine++
	}

	board := Board{lines: lines}
	for _, xCoord := range xCoords {
		iAnswer += FindXmases(xCoord, board)
	}

	fmt.Println(iAnswer)
}

func Day4Puzzle2() {

	file, err := os.Open("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	reDos := regexp.MustCompile(`A`)
	scanner := bufio.NewScanner(file)
	iAnswer := 0
	iLine := 0
	aCoords := make([]XCoord, 0)
	lines := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		As := reDos.FindAllStringIndex(s, -1)
		for _, A := range As {
			aCoord := XCoord{x: A[0], y: iLine}
			aCoords = append(aCoords, aCoord)
		}
		lines = append(lines, s)
		iLine++
	}

	board := Board{lines: lines}
	for _, aCoord := range aCoords {
		iAnswer += FindXXmases(aCoord, board)
	}

	fmt.Println(iAnswer)
}

func FindXXmases(aCoord XCoord, board Board) int {
	if aCoord.x == 0 || aCoord.x == len(board.lines[0])-1 {
		return 0
	}
	if aCoord.y == 0 || aCoord.y == len(board.lines)-1 {
		return 0
	}
	sUpLeft := board.LetterAt(aCoord.x-1, aCoord.y-1)
	sUpRight := board.LetterAt(aCoord.x+1, aCoord.y-1)
	sDownRight := board.LetterAt(aCoord.x+1, aCoord.y+1)
	sDownLeft := board.LetterAt(aCoord.x-1, aCoord.y+1)

	sLoop := sUpLeft + sUpRight + sDownRight + sDownLeft

	idx := strings.Index(sLoop, "MM")

	switch idx {
	case 0:
		if sLoop[2] == 'S' && sLoop[3] == 'S' {
			return 1
		}
	case 1:
		if sLoop[0] == 'S' && sLoop[3] == 'S' {
			return 1
		}
	case 2:
		if sLoop[0] == 'S' && sLoop[1] == 'S' {
			return 1
		}
	}

	if sLoop[0] == 'M' && sLoop[3] == 'M' && sLoop[1] == 'S' && sLoop[2] == 'S' {
		return 1
	}

	return 0
}

func FindXmases(xCoord XCoord, board Board) int {
	iFound := 0
	// Look up
	if board.Is("M", xCoord.x, xCoord.y-1) {
		if board.Is("A", xCoord.x, xCoord.y-2) {
			if board.Is("S", xCoord.x, xCoord.y-3) {
				iFound++
			}
		}

	}
	// Look down
	if board.Is("M", xCoord.x, xCoord.y+1) {
		if board.Is("A", xCoord.x, xCoord.y+2) {
			if board.Is("S", xCoord.x, xCoord.y+3) {
				iFound++
			}
		}
	}
	// Look left
	if board.Is("M", xCoord.x-1, xCoord.y) {
		if board.Is("A", xCoord.x-2, xCoord.y) {
			if board.Is("S", xCoord.x-3, xCoord.y) {
				iFound++
			}
		}
	}
	// Look right
	if board.Is("M", xCoord.x+1, xCoord.y) {
		if board.Is("A", xCoord.x+2, xCoord.y) {
			if board.Is("S", xCoord.x+3, xCoord.y) {
				iFound++
			}
		}
	}
	// Look up-left
	if board.Is("M", xCoord.x-1, xCoord.y-1) {
		if board.Is("A", xCoord.x-2, xCoord.y-2) {
			if board.Is("S", xCoord.x-3, xCoord.y-3) {
				iFound++
			}
		}
	}
	// Look up-right
	if board.Is("M", xCoord.x+1, xCoord.y-1) {
		if board.Is("A", xCoord.x+2, xCoord.y-2) {
			if board.Is("S", xCoord.x+3, xCoord.y-3) {
				iFound++
			}
		}
	}
	// Look down-left
	if board.Is("M", xCoord.x-1, xCoord.y+1) {
		if board.Is("A", xCoord.x-2, xCoord.y+2) {
			if board.Is("S", xCoord.x-3, xCoord.y+3) {
				iFound++
			}
		}
	}
	// Look down-right
	if board.Is("M", xCoord.x+1, xCoord.y+1) {
		if board.Is("A", xCoord.x+2, xCoord.y+2) {
			if board.Is("S", xCoord.x+3, xCoord.y+3) {
				iFound++
			}
		}
	}
	return iFound
}

func (b *Board) Is(letter string, x int, y int) bool {
	if x < 0 || y < 0 || y >= len(b.lines) || x >= len(b.lines[y]) {
		return false
	}
	return string(b.lines[y][x]) == letter
}

func (b *Board) LetterAt(x int, y int) string {
	if x < 0 || y < 0 || y >= len(b.lines) || x >= len(b.lines[y]) {
		return ""
	}
	return string(b.lines[y][x])
}
