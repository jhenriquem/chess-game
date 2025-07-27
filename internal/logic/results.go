package logic

import (
	"strings"
)

// Functions responsible for verifying the results of movements (check, checkmate)

func IsCheck(result rune) bool {
	if result == '+' {
		return true
	}
	return false
}

func IsMate(result rune) bool {
	if result == '#' {
		return true
	}
	return false
}

func VerifyResultsOfMoves(gameString string) string {
	str := strings.ReplaceAll(gameString, " ", "")
	resultRune := rune(str[len(str)-2])

	if check := IsCheck(resultRune); check {
		return "Is check"
	} else if mate := IsMate(resultRune); mate {
		return "Is checkmate"
	}

	return ""
}
