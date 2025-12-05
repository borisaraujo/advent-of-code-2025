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

func main() {

	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	// comparison_grid := make([]string, len(grid))
	total := 0
	for i, line := range grid {
		for j, c := range line {

			if ((i == 0 && j == 0) ||
				(i == 0 && j == len(line)-1) ||
				(i == len(grid)-1 && j == 0) ||
				(i == len(grid)-1 && j == len(line)-1)) &&
				c == '@' {
				total++
			} else if i == 0 {
				if c == '@' {
					chars := string(grid[i][j-1]) +
						string(grid[i][j+1]) +
						string(grid[i+1][j]) +
						string(grid[i+1][j+1]) +
						string(grid[i+1][j-1])
					if countAtSigns(chars) < 4 {
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
						total++
					}
				}
			}

		}
	}

	fmt.Println(total)
}
