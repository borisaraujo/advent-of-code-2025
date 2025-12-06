package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	worksheet := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Join(strings.Fields(line), " ")
		values := strings.Split(line, " ")
		wsLine := make([]string, 0)
		wsLine = append(wsLine, values...)
		worksheet = append(worksheet, wsLine)
	}

	total := 0
	for j := range len(worksheet[0]) {
		operator := worksheet[len(worksheet)-1][j]
		var colTotal int
		if operator == "+" {
			colTotal = 0
		} else if operator == "*" {
			colTotal = 1
		}
		for i := 0; i < len(worksheet)-1; i++ {
			val, _ := strconv.Atoi(worksheet[i][j])
			if operator == "+" {
				colTotal += val
			} else if operator == "*" {
				colTotal *= val
			}
		}
		total += colTotal
	}
	println(total)
}
