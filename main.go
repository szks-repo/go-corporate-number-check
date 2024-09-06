package main

import (
	"regexp"
	"strconv"
	"strings"
)

var corpNumPattern = regexp.MustCompile(`^[1-9]{1}\d{12}$`)

func CheckCorporateNumber(number string) bool {
	if !corpNumPattern.MatchString(number) {
		return false
	}

	checkDigit, _ := strconv.Atoi(number[0:1])
	baseNums := make([]int, 12)
	for i, s := range strings.Split(number[1:], "") {
		baseNum, _ := strconv.Atoi(s)
		baseNums[i] = baseNum
	}

	var total int
	for n := 1; n <= len(baseNums); n++ {
		Q := 1
		if n%2 == 0 {
			Q = 2
		}

		P := baseNums[len(baseNums)-n]
		total += Q * P
	}
	gotCheckDigit := 9 - (total % 9)
	return gotCheckDigit == checkDigit
}
