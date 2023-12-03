package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SchematicRecord []rune
type Schematic []SchematicRecord

func Day3Puzzle1() {
	schematic := parseSchematic()

	symbolMap := [][]bool{}
	intMap := [][]int{}

	for idxRow, schematicRecord := range schematic {
		symbolMap = append(symbolMap, make([]bool, len(schematicRecord)))

		// Initialize a slice of ints with -1
		intSlice := make([]int, len(schematicRecord))
		for i := range intSlice {
			intSlice[i] = -1
		}

		// Append the initialized int slice
		intMap = append(intMap, intSlice)

		for idxCol, item := range schematicRecord {
			if item != '.' {
				iVal := int(item - '0')
				if iVal >= 0 && iVal <= 9 {
					intMap[idxRow][idxCol] = iVal
				} else {
					symbolMap[idxRow][idxCol] = true
				}
			}
		}
	}

	parts := findParts(intMap, symbolMap)
	//fmt.Println(parts)
	iSumParts := 0
	for _, iPart := range parts {
		iSumParts += iPart
	}
	fmt.Println("Answer:", iSumParts)
}

func Day2Puzzle2() {
	iAnswer := 0

	fmt.Println("Answer: ", iAnswer)
}

func parseSchematic() Schematic {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	schematic := []SchematicRecord{}
	iRecordCounter := 0
	for scanner.Scan() {
		iRecordCounter++
		s := scanner.Text()

		schematicRecord := SchematicRecord{}
		for _, sItem := range s {
			schematicRecord = append(schematicRecord, sItem)
		}
		schematic = append(schematic, schematicRecord)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return schematic
}

func findParts(intMap [][]int, symbolMap [][]bool) []int {

	parts := []int{}
	for idxRow, row := range intMap {
		for i := 0; i < len(row); i++ {
			idxStartCol := -1
			partCandidate := row[i]
			if partCandidate > -1 {
				idxStartCol = i
				i++
				for i < len(row) && row[i] > -1 {
					partCandidate = (partCandidate * 10) + row[i]
					i++
				}
				bIsPart := isPart(symbolMap, partCandidate, idxRow, idxStartCol)
				if bIsPart {
					parts = append(parts, partCandidate)
				}
			}
		}
	}
	return parts
}

func isPart(symbolMap [][]bool, partCanidate int, idxRow int, idxStartCol int) bool {

	sPartCandidate := strconv.Itoa(partCanidate)
	iPartLength := len(sPartCandidate)

	idxColToLeft := idxStartCol - 1
	idxColToRight := idxStartCol + iPartLength
	idxRowAbove := idxRow - 1
	idxRowBelow := idxRow + 1

	// look for symbol to left
	if idxColToLeft > -1 && symbolMap[idxRow][idxColToLeft] {
		return true
	}
	// look for symbol to right
	if idxColToRight < len(symbolMap[idxRow]) &&
		symbolMap[idxRow][idxColToRight] {
		return true
	}
	// look for symbol above
	if idxRowAbove >= 0 {
		for i := max(idxColToLeft, 0); i <= min(idxColToRight, len(symbolMap[idxRow])-1); i++ {
			if symbolMap[idxRowAbove][i] {
				return true
			}
		}
	}

	// look for symbol below
	if idxRowBelow < len(symbolMap) {
		for i := max(idxColToLeft, 0); i <= min(idxColToRight, len(symbolMap[idxRow])-1); i++ {
			if symbolMap[idxRowBelow][i] {
				return true
			}
		}
	}

	return false
}
