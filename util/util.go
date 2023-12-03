package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetIntFromString(s string) int {
	iCur, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return iCur
}

func GetIntFromEnglish(s string) int {
	sLow := strings.ToLower(s)
	i := 0
	switch sLow {
	case "one":
		i = 1
	case "two":
		i = 2
	case "three":
		i = 3
	case "four":
		i = 4
	case "five":
		i = 5
	case "six":
		i = 6
	case "seven":
		i = 7
	case "eight":
		i = 8
	case "nine":
		i = 9
	default:
		fmt.Println("*************", sLow)
		os.Exit(2)
	}
	if i < 1 || i > 9 {
		os.Exit(2)
	}
	return i
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
