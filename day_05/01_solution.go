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

func processNonDashLine(line string, h *[]int, rangeMap map[int]int, debug bool) bool {
	// process integer
	n, err := strconv.Atoi(line)

	if err != nil {
		return false
	}

	if debug {
		fmt.Printf("\nChecking number: %d\n", n)
	}

	for _, rangeMax := range *h {
		if n <= rangeMax {
			rangeMin := rangeMap[rangeMax]
			if debug {
				fmt.Printf("Checking if %d is between %d and %d: ", n, rangeMin, rangeMax)
			}

			if rangeMin <= n {
				if debug {
					fmt.Printf("true\n%d is fresh\n", n)
				}
				return true
			}

			if debug {
				fmt.Printf("false\n")
			}
		} else {
			if debug {
				fmt.Printf("%d is spoiled\n", n)
			}
			return false
		}
	}

	if debug {
		fmt.Printf("%d is spoiled\n", n)
	}
	return false
}

func processDashedLine(line string, maxSlice *[]int, maxMinMap map[int]int) {
	// split - between range
	idRangeMin, idRangeMax := processRangeString(line)

	// add to heap and hashmap
	pushSorted(maxSlice, idRangeMax)
	addKV(maxMinMap, idRangeMax, idRangeMin)
}

func main() {
	content := readfile("input.txt")
	contentLines := splitLines(content)

	// data structures
	freshIdMaxToMin := &[]int{}
	freshIdMaxMinMap := map[int]int{}

	res := 0

	for _, line := range contentLines {
		if strings.Contains(line, "-") {
			processDashedLine(line, freshIdMaxToMin, freshIdMaxMinMap)
			// continue to next line
			continue
		}

		if processNonDashLine(line, freshIdMaxToMin, freshIdMaxMinMap, false) {
			res++
		}
	}

	fmt.Println(res)
}
