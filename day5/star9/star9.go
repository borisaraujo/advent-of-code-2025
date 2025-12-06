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
	var ranges, ids []string
	flagRanges := true
	fresh := make(map[int]bool)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			flagRanges = false
			continue
		}
		if flagRanges {
			ranges = append(ranges, line)
		} else {
			ids = append(ids, line)
		}
	}

	for _, rng := range ranges {
		parts := strings.Split(rng, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for _, id := range ids {
			i, _ := strconv.Atoi(id)
			if start <= i && i <= end {
				fresh[i] = true
			}
		}

	}

	for _, id := range ids {
		i, _ := strconv.Atoi(id)
		_, exists := fresh[i]

		if exists {
			total += 1
		}
	}

	println(total)
}
