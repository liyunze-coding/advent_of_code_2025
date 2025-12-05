package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func splitLines(s string) []string {
	lines := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

// check if ID falls below the MAX IDs
// check if unknown ID falls between in all the max IDs it matches

func processRangeString(s string) (int, int) {
	idRange := strings.Split(s, "-")

	idRangeMinString := idRange[0]
	idRangeMaxString := idRange[1]

	idRangeMin, err := strconv.Atoi(idRangeMinString)

	if err != nil {
		log.Fatal(err)
	}

	idRangeMax, err := strconv.Atoi(idRangeMaxString)

	if err != nil {
		log.Fatal(err)
	}

	return idRangeMin, idRangeMax
}

// push sorted from max to min
func pushSorted(h *[]int, n int) {
	if slices.Contains(*h, n) {
		return
	}
	for i, x := range *h {
		if n > x {
			*h = slices.Insert(*h, i, n)
			return
		}
	}

	*h = append(*h, n)
}

func addKV(m map[int]int, key int, value int) {
	mapValue, ok := m[key]

	if !ok {
		m[key] = value
	} else if mapValue > value {
		m[key] = value
	}
}

func processDashedLine(line string, maxSlice *[]int, maxMinMap map[int]int) {
	// split - between range
	idRangeMin, idRangeMax := processRangeString(line)

	// add to heap and hashmap
	pushSorted(maxSlice, idRangeMax)
	addKV(maxMinMap, idRangeMax, idRangeMin)
}

func processRanges(freshIdMaxToMin *[]int, rangeMap map[int]int) int {
	res := 0
	maxRanges := *freshIdMaxToMin
	currentMax := maxRanges[0]
	currentMin := rangeMap[currentMax]
	rangeLen := len(maxRanges)

	// IF MAX ARE EQUAL, TAKE LOWER MIN -- NEVERMIND NOT POSSIBLE ANYMORE
	// IF MAX BELOW BUT MIN IS HIGHER OR EQUAL THAN CURRENT MIN, IGNORE
	// IF MAX BELOW AND MIN LOWER THAN CURRENT MIN:
	// MAX ABOVE THE CURRENT MIN: overlapped
	// MAX BELOW CURRENT MIN: count previous overlapped, move on to next
	for i := 1; i < rangeLen; i++ {

		tempMax := maxRanges[i]
		tempMin := rangeMap[tempMax]

		if currentMax > tempMax && tempMin < currentMin {
			if tempMax >= currentMin {
				// overlap
				currentMin = tempMin
			} else {
				// process previous overlap
				// fmt.Println("else", currentMax, currentMin)
				res += currentMax - currentMin + 1

				// reassign
				currentMax = tempMax
				currentMin = tempMin
			}
		}
	}

	// fmt.Println("last", currentMax, currentMin)
	res += currentMax - currentMin + 1

	return res
}

func main() {
	content := readfile("input.txt")
	contentLines := splitLines(content)

	// data structures
	freshIdMaxToMin := &[]int{}
	rangeMap := map[int]int{}
	res := 0

	for _, line := range contentLines {
		if strings.Contains(line, "-") {
			processDashedLine(line, freshIdMaxToMin, rangeMap)
			// continue to next line
			continue
		}

		// finished processing dashed lines
		if line == "" {
			// process the ranges
			res = processRanges(freshIdMaxToMin, rangeMap)
			break
		}
	}

	fmt.Println(res)
}
