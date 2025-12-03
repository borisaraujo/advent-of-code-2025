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

			for j := 1; j < len(iStr); j++ {

				if len(iStr)%j != 0 {
					continue
				}
				repeatedSequence := true
				startingElement := iStr[:j]

				for k := 0; k < len(iStr); k += j {
					if iStr[k:k+j] != startingElement {
						repeatedSequence = false
						break
					}
				}
				if repeatedSequence {
					counter += i
					break
				}
			}

		}
	}
	fmt.Println(counter)
}
