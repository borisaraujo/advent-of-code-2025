package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items [12]int
	top   int
}

func (s *Stack) pop() (int, error) {
	if s.top > 0 {
		s.top--
		return s.items[s.top], nil
	}
	return -1, errors.New("Stack underflow")
}

func (s *Stack) push(item int) error {
	if s.top < 12 {
		s.items[s.top] = item
		s.top++
		return nil
	}
	return errors.New("Stack overflow")
}

func (s *Stack) isFull() bool {
	return s.top == 12
}

func (s *Stack) spaceLeft() int {
	return 12 - s.top
}

func (s *Stack) size() int {
	return s.top
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func intSliceToInt(intSlice []int) int {
	var b strings.Builder

	for _, v := range intSlice {
		b.WriteString(strconv.Itoa(v))
	}
	s := b.String()
	x, _ := strconv.Atoi(s)
	return x
}

func findJoltage(powerBankInt []int) int {
	var stack Stack

	powerBankSize := len(powerBankInt)
	for i, n := range powerBankInt {
		if stack.isEmpty() {
			stack.push(n)
			continue
		}

		itemsLeftPB := powerBankSize - i - 1
		nUsed := false

		var aux Stack
		for stack.spaceLeft() <= itemsLeftPB && stack.size() != 0 {
			top, _ := stack.pop()
			aux.push(top)
		}

		for !aux.isEmpty() {
			top, _ := aux.pop()

			if n > top {
				stack.push(n)
				nUsed = true
				break
			} else {
				stack.push(top)
			}
		}

		if !nUsed && !stack.isFull() {
			stack.push(n)
		}
	}

	total := intSliceToInt(stack.items[:])
	return total
}

func main() {
	file, _ := os.Open("../input.txt")

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		powerBank := scanner.Text()
		powerBankInt := make([]int, len(powerBank))
		for i, c := range powerBank {
			powerBankInt[i] = int(c - '0')
		}
		total += findJoltage(powerBankInt)
	}

	fmt.Println(total)
}
