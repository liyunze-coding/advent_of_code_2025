package main

import (
	"fmt"
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

func main() {
	content := readfile("input.txt")

	idSum := 0

	ids := strings.SplitSeq(content, ",")

	// split each element from id1-id2 to {id1 id2}
	for idRange := range ids {
		numString := strings.Split(idRange, "-")
		num1Int, err := strconv.Atoi(numString[0])

		if err != nil {
			panic(err)
		}

		num2Int, err := strconv.Atoi(numString[1])

		// WE HAVE FIRST NUMBER AND SECOND NUMBER.
		// DO A FOR LOOP BETWEEN THEM AND CHECK IF FIRST HALF == SECOND HALF

		for i := num1Int; i <= num2Int; i++ {
			iStr := strconv.Itoa(i)

			strLenHalf := len(iStr) / 2
			firstHalf := iStr[:strLenHalf]
			secondHalf := iStr[strLenHalf:]

			// fmt.Printf("%s %s\n", firstHalf, secondHalf)

			if firstHalf == secondHalf {
				idSum += i
			}
		}
	}

	fmt.Println(idSum)
}
