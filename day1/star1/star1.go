package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	value    int
	previous *Number
	next     *Number
}

type Dial struct {
	zeroClicks    int
	currentNumber *Number
}

func (d *Dial) turnLeft(number int) {
	for range number {
		d.currentNumber = d.currentNumber.previous
	}
	if d.currentNumber.value == 0 {
		d.zeroClicks++
	}
}

func (d *Dial) turnRight(number int) {
	for range number {
		d.currentNumber = d.currentNumber.next
	}
	if d.currentNumber.value == 0 {
		d.zeroClicks++
	}
}

func (d *Dial) turn(instruction string) {
	var direction rune
	var number int
	direction = rune(instruction[0])
	number, _ = strconv.Atoi(instruction[1:])
	switch direction {
	case 'L':
		d.turnLeft(number)
	case 'R':
		d.turnRight(number)
	}
}

func createDial() Dial {
	var zero *Number = &Number{0, nil, nil}
	var currentNumber, newNumber *Number
	var dial Dial
	currentNumber = zero
	for i := 1; i < 100; i++ {
		newNumber = &Number{i, currentNumber, nil}
		currentNumber.next = newNumber
		currentNumber = newNumber
		if i == 50 {
			dial.currentNumber = currentNumber
		}
	}
	currentNumber.next = zero
	zero.previous = currentNumber
	currentNumber = zero
	return dial
}

func main() {

	var dial Dial = createDial()
	file, _ := os.Open("../input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instruction := scanner.Text()
		dial.turn(instruction)
	}
	fmt.Println(dial.zeroClicks)
}
