package main

import (
	"bufio"
	"os"
	"testing"
)

func TestCheckCorporateNumber(t *testing.T) {
	t.Parallel()

	for _, number := range getTestNumbers(t) {
		t.Run(number, func(t *testing.T) {
			t.Parallel()
			if !CheckCorporateNumber(number) {
				t.Errorf("want true, got false")
			}
		})
	}
}

func getTestNumbers(t *testing.T) []string {
	t.Helper()

	f, err := os.Open("corpNums.txt")
	if err != nil {
		t.Fatal(err)
	}

	var numbers []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if txt := scanner.Text(); txt != "" {
			numbers = append(numbers, txt)
		}
	}

	if len(numbers) == 0 {
		t.Fatal("no numbers found")
	}

	return numbers
}
