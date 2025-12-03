package main

import (
	"fmt"
	"iter"
	"log"
	"os"
	"strconv"
	"strings"
)

func readfile(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func splitLines(s string) iter.Seq[string] {
	return strings.SplitSeq(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}

func processLine(line string) (int, int) {
	largestDigit := -1
	largestDigitIndex := -1

	stringSlice := strings.Split(line, "")

	stringSliceLen := len(stringSlice)

	// from last 2nd to 1st digit, find largest digit
	for i := stringSliceLen - 2; i >= 0; i-- {
		digit, err := strconv.Atoi(stringSlice[i])

		if err != nil {
			log.Fatal(err)
		}

		// if digits equal but more at the front, replace index
		if digit >= largestDigit {
			largestDigit = digit
			largestDigitIndex = i
		}
	}

	secondDigit := -1

	// from last digit to largest digit, find second
	for j := largestDigitIndex + 1; j < stringSliceLen; j++ {
		digit, err := strconv.Atoi(stringSlice[j])

		if err != nil {
			log.Fatal(err)
		}

		if digit >= secondDigit {
			secondDigit = digit
		}
	}

	fmt.Printf("%s: %d %d\n", line, largestDigit, secondDigit)
	return largestDigit, secondDigit
}

func main() {
	content := readfile("input.txt")

	lines := splitLines(content)

	totalJoltage := 0

	for line := range lines {
		num1, num2 := processLine(line)

		if num1 != -1 {
			totalJoltage += num1*10 + num2
		}

	}

	fmt.Println(totalJoltage)
}
