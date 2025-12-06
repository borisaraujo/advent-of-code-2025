package main

import (
	"bufio"
	"fmt"
	"os"
)

func countAtSigns(grid []string, i, j int) int {
	total := 0
	rows, cols := len(grid), len(grid[0])
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if (x == i && y == j) || (x == -1 || y == -1 || x == rows || y == cols) {
				continue
			}
			if grid[x][y] == '@' {
				total++
			}
		}
	}
	return total
}

func main() {

	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	total := 0
	for i, line := range grid {
		for j, c := range line {
			if c == '@' && countAtSigns(grid, i, j) < 4 {
				total++
			}
		}
	}

	fmt.Println(total)
}
