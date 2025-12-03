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

func all(strings []string) bool {
	firstS := strings[0]

	for i, s := range strings {
		if i == 0 {
			continue
		}

		if firstS != s {
			return false
		}
	}

	return true
}

func stringToArray(s string, size int) []string {
	var res []string

	var j int
	for i := 0; i < len(s); i += size {
		j += size
		if j > len(s) {
			j = len(s)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		// fmt.Println(s[i:j])
		res = append(res, s[i:j])
	}

	return res
}

func hasRepeat(s string) bool {
	for i := 1; i < len(s)/2+1; i++ {
		stringArray := stringToArray(s, i)

		if all(stringArray) {
			return true
		}
	}

	return false
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

			if hasRepeat(iStr) {
				idSum += i
			}
		}
	}

	fmt.Println(idSum)
}
