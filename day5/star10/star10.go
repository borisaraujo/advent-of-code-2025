package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func mergeRanges(aRange *Range, bRange *Range) *Range {
	newRange := &Range{}
	if aRange.end < bRange.start || bRange.end < aRange.start {
		return nil
	}
	newRange.start = min(aRange.start, bRange.start)
	newRange.end = max(aRange.end, bRange.end)
	return newRange
}

func mergeRangeSlice(ranges []Range) []Range {
	newRanges := make([]Range, 0)
	rangeMap := make(map[Range]bool, 0)

	for i, iVal := range ranges {
		_, iExists := rangeMap[iVal]
		if iExists {
			continue
		}
		for j := i + 1; j < len(ranges); j++ {
			_, jExists := rangeMap[ranges[j]]
			if jExists {
				continue
			}
			newRange := mergeRanges(&iVal, &ranges[j])
			if newRange != nil {
				newRanges = append(newRanges, *newRange)
				rangeMap[iVal] = true
				rangeMap[ranges[j]] = true
				continue
			}
		}
	}

	mergedRanges := make([]Range, 0)
	for _, val := range ranges {
		_, exists := rangeMap[val]

		if exists {
			continue
		}

		mergedRanges = append(mergedRanges, val)
	}

	mergedRanges = append(mergedRanges, newRanges...)

	return mergedRanges
}

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	var ranges []Range
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		newRange := Range{start: start, end: end}
		ranges = append(ranges, newRange)
	}

	newRanges := make([]Range, 0)
	for len(ranges) != len(newRanges) {
		if len(newRanges) != 0 {
			ranges = newRanges
		}
		newRanges = mergeRangeSlice(ranges)
	}

	for _, rng := range newRanges {
		total += rng.end - rng.start + 1
	}
	fmt.Println(total)
}
