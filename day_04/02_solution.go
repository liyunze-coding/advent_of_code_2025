package main

import (
	"fmt"
	"os"
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

func checkAdjacent(grid [][]rune, mainX int, mainY int, width int, height int) bool {
	xMin := max(mainX-1, 0)
	yMin := max(mainY-1, 0)
	xMax := min(mainX+1, width-1)
	yMax := min(mainY+1, height-1)

	count := 0

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if y == mainY && x == mainX {
				continue
			}
			if grid[y][x] == '@' {
				count++
			}
		}
	}

	return count < 4
}

func processGrid(grid [][]rune, gridWidth int, gridHeight int) (int, [][]rune) {
	removedCount := 0
	gridCopy := grid

	for y, row := range grid {
		for x, char := range row {
			if char == '@' && checkAdjacent(grid, x, y, gridWidth, gridHeight) {
				removedCount++
				gridCopy[y][x] = '.'
			}
		}
	}

	return removedCount, gridCopy
}

func main() {
	content := readfile("input.txt")
	contentLines := splitLines(content)

	// FILE CONTENT TO GRID OF CHAR STRINGS
	grid := [][]rune{}

	gridWidth := -1
	gridHeight := len(contentLines)

	for _, line := range contentLines {
		gridLine := []rune{}
		for _, char := range line {
			gridLine = append(gridLine, char)
		}
		grid = append(grid, gridLine)

		// length of row
		if gridWidth == -1 {
			gridWidth = len(gridLine)
		}
	}

	res := 0
	removedCount := -1

	// LOOP OVER EACH GRID ELEMENT
	for removedCount != 0 {
		removedCountTemp, newGrid := processGrid(grid, gridWidth, gridHeight)
		grid = newGrid
		res += removedCountTemp
		removedCount = removedCountTemp
	}

	fmt.Println(res)
}
