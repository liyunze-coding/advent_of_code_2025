package main

import (
	"fmt"
	"iter"
	"log"
	"math"
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

func processLine(line string) int {
	largestDigit := -1
	largestDigitIndex := -1

	stringSlice := strings.Split(line, "")

	stringSliceLen := len(stringSlice)

	// from last 12th to 1st digit, find largest digit
	for i := stringSliceLen - 12; i >= 0; i-- {
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

	currentLargestIdx := largestDigitIndex
	numberSlice := []int{largestDigit}

	for j := 11; j > 0; j-- {
		largestDigitTemp := -1
		largestDigitIdxTemp := -1

		for k := currentLargestIdx + 1; k < stringSliceLen-j+1; k++ {
			digit, err := strconv.Atoi(stringSlice[k])

			if err != nil {
				log.Fatal(err)
			}

			if digit > largestDigitTemp {
				largestDigitTemp = digit
				largestDigitIdxTemp = k
			}
		}

		currentLargestIdx = largestDigitIdxTemp
		numberSlice = append(numberSlice, largestDigitTemp)
	}

	// return it as a full number
	res := 0
	for x := range 12 {
		res += int(math.Pow(10, float64(x))) * numberSlice[11-x]
	}

	return res
}

func main() {
	content := readfile("input.txt")

	lines := splitLines(content)

	totalJoltage := 0

	for line := range lines {
		totalJoltage += processLine(line)
	}

	fmt.Println(totalJoltage)
}
