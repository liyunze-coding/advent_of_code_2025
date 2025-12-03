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

	for i := range lines {
		word := lines[i]

		if word == "" {
			break
		}

		direction := word[0:1]
		distance, err := strconv.Atoi(word[1:])

		if err != nil {
			log.Fatal(err)
		}

		// TURN THE DIAL
		switch direction {
		case "L":
			position = (position - distance) % 100

			for position < 0 {
				position += 100
			}
		case "R":
			position = (position + distance) % 100
		}

		if position == 0 {
			zeros++
		}

	}

	fmt.Printf("%d", zeros)
}
