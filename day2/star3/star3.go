package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ids string
	file, _ := os.Open("../input.txt")
	defer file.Close()

	counter := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ids = scanner.Text()
	}

	ranges := strings.Split(ids, ",")

	for _, idRange := range ranges {
		rangeLimits := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(rangeLimits[0])
		end, _ := strconv.Atoi(rangeLimits[1])

		for i := start; i <= end; i++ {
			iStr := strconv.Itoa(i)
			if len(iStr)%2 != 0 {
				continue
			}
			firstHalf := iStr[:len(iStr)/2]
			secondHalf := iStr[len(iStr)/2:]

			if firstHalf == secondHalf {
				counter += i
			}
		}
	}
	fmt.Println(counter)
}
