package main

import (
	"fmt"
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

func main() {
	content := readfile("input.txt")
	position := 50
	zeros := 0

	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")

	fmt.Printf("%d\n", position)

	for i := range lines {
		word := lines[i]

		if word == "" {
			break
		}

		direction := word[0:1] // L or R
		distance, err := strconv.Atoi(word[1:])

		if err != nil {
			log.Fatal(err)
		}

		// TURN THE DIAL
		switch direction {
		case "L":
			beforePosition := position
			position = position - distance

			for position < 0 {
				position += 100
				zeros++
			}

			if beforePosition == 0 {
				zeros--
			}

			if position == 0 {
				zeros++
			}
		case "R":
			position = position + distance

			for position >= 100 {
				position -= 100
				zeros++
			}
		}

		fmt.Printf("%s %d, %d\n", word, position, zeros)
	}

	fmt.Printf("\n%d\n", zeros)
}
