package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getColsSizes(row string) []int {
	colSize := 0
	var colsSizes []int
	for _, c := range row {
		if (c == '+' || c == '*') && colSize != 0 {
			colsSizes = append(colsSizes, colSize-1)
			colSize = 0
		}
		colSize++
	}
	colsSizes = append(colsSizes, colSize)
	return colsSizes
}

func StringsToBytes(ss []string) [][]byte {
	out := make([][]byte, len(ss))
	for i, s := range ss {
		out[i] = []byte(s)
	}
	return out
}

func BytesToStrings(bb [][]byte) []string {
	out := make([]string, len(bb))
	for i := range bb {
		out[i] = string(bb[i])
	}
	return out
}

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	worksheet := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		worksheet = append(worksheet, line)
	}

	lastRow := worksheet[len(worksheet)-1]

	colsSizes := getColsSizes(lastRow)

	worksheetBytes := StringsToBytes(worksheet)

	for i, row := range worksheet {
		if i == len(worksheet)-1 {
			break
		}
		col := 0
		colSize := 0
		for j, c := range row {
			if colSize == colsSizes[col] {
				colSize = 0
				col++
				continue
			}
			if c == ' ' && colSize < colsSizes[col] {
				worksheetBytes[i][j] = '@'
			}
			colSize++
		}
	}

	worksheet = BytesToStrings(worksheetBytes)
	newWorksheet := make([][]string, 0)
	for _, line := range worksheet {
		newLine := strings.Join(strings.Fields(line), " ")
		values := strings.Split(newLine, " ")
		wsLine := make([]string, 0)
		wsLine = append(wsLine, values...)
		newWorksheet = append(newWorksheet, wsLine)
	}

	total := 0

	for j := range len(newWorksheet[0]) {
		operator := newWorksheet[len(newWorksheet)-1][j]
		var colTotal int
		switch operator {
		case "+":
			colTotal = 0
		case "*":
			colTotal = 1
		}
		maxLen := len(newWorksheet[0][j])
		newArr := make([]string, maxLen)

		for i := 0; i < len(newWorksheet)-1; i++ {
			subCols := strings.Split(newWorksheet[i][j], "")
			for k := range newArr {
				if subCols[k] != "@" {
					newArr[k] += subCols[k]
				}
			}
		}

		for _, valStr := range newArr {
			valInt, _ := strconv.Atoi(valStr)
			switch operator {
			case "+":
				colTotal += valInt
			case "*":
				colTotal *= valInt
			}
		}

		total += colTotal

	}

	fmt.Println(total)

}
