package main

import (
	"bufio"
	"fmt"
	"os"
)

func findTwoGreatest(powerBankInt []int) int {
	var greatest int
	for i := 0; i < len(powerBankInt)+1; i++ {
		for j := i + 1; j < len(powerBankInt); j++ {
			total := 10*powerBankInt[i] + powerBankInt[j]
			if total > greatest {
				greatest = total
			}
		}
	}
	return greatest
}

func main() {
	var total int
	file, _ := os.Open("../input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		powerBank := scanner.Text()
		powerBankInt := make([]int, len(powerBank))

		for i, c := range powerBank {
			powerBankInt[i] = int(c - '0')

		}
		total += findTwoGreatest(powerBankInt)

	}
	fmt.Println(total)
}
