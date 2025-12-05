package main

import (
	"bufio"
	"fmt"
	"os"
)

func countAtSigns(charArray string) int {
	total := 0

	for _, c := range charArray {
		if c == '@' {
			total++
		}
	}
	return total
}

func stringsToRunes(grid []string) [][]rune {
	out := make([][]rune, len(grid))
	for i, line := range grid {
		out[i] = []rune(line)
	}
	return out
}

func runesToStrings(grid [][]rune) []string {
	out := make([]string, len(grid))
	for i, line := range grid {
		out[i] = string(line)
	}
	return out
}

func processPaper(grid []string) ([]string, int) {
	total := 0
	controlGrid := stringsToRunes(grid)
	for i, line := range grid {
		for j, c := range line {
			controlGrid[i][j] = c
			if ((i == 0 && j == 0) ||
				(i == 0 && j == len(line)-1) ||
				(i == len(grid)-1 && j == 0) ||
				(i == len(grid)-1 && j == len(line)-1)) &&
				c == '@' {
				controlGrid[i][j] = '.'
				total++
			} else if i == 0 {
				if c == '@' {
					chars := string(grid[i][j-1]) +
						string(grid[i][j+1]) +
						string(grid[i+1][j]) +
						string(grid[i+1][j+1]) +
						string(grid[i+1][j-1])
					if countAtSigns(chars) < 4 {
						controlGrid[i][j] = '.'
						total++
					}
				}
			} else if i == len(grid)-1 {
				if c == '@' {
					chars := string(grid[i][j-1]) +
						string(grid[i][j+1]) +
						string(grid[i-1][j]) +
						string(grid[i-1][j+1]) +
						string(grid[i-1][j-1])
					if countAtSigns(chars) < 4 {
						controlGrid[i][j] = '.'
						total++
					}
				}
			} else if j == 0 {
				if c == '@' {
					chars := string(grid[i+1][j]) +
						string(grid[i-1][j]) +
						string(grid[i-1][j+1]) +
						string(grid[i][j+1]) +
						string(grid[i+1][j+1])
					if countAtSigns(chars) < 4 {
						controlGrid[i][j] = '.'
						total++
					}
				}
			} else if j == len(line)-1 {
				if c == '@' {
					chars := string(grid[i+1][j]) +
						string(grid[i-1][j]) +
						string(grid[i-1][j-1]) +
						string(grid[i][j-1]) +
						string(grid[i+1][j-1])
					if countAtSigns(chars) < 4 {
						controlGrid[i][j] = '.'
						total++
					}
				}
			} else {
				if c == '@' {
					chars := string(grid[i+1][j]) +
						string(grid[i-1][j]) +
						string(grid[i-1][j-1]) +
						string(grid[i][j-1]) +
						string(grid[i+1][j-1]) +
						string(grid[i-1][j+1]) +
						string(grid[i][j+1]) +
						string(grid[i+1][j+1])
					if countAtSigns(chars) < 4 {
						controlGrid[i][j] = '.'
						total++
					}
				}
			}

		}
	}

	return runesToStrings(controlGrid), total
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

	for {
		newGrid, newTotal := processPaper(grid)
		grid = newGrid
		total += newTotal
		if newTotal == 0 {
			break
		}
	}

	fmt.Println(total)
}
