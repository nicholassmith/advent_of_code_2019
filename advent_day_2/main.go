package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var positions []int

	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	for _, c := range strings.Split(string(dat), ",") {
		x, err := strconv.Atoi(strings.TrimSpace(string(c)))
		check(err)
		positions = append(positions, x)
	}

	fmt.Printf("part 1: %d\n", compute(positions))
}

func compute(positions []int) int {
	for i := 0; i < len(positions); i += 4 {
		opcode := positions[i]

		switch opcode {
		case 1:
			res := positions[positions[i+1]] + positions[positions[i+2]]
			positions[positions[i+3]] = res
		case 2:
			res := positions[positions[i+1]] * positions[positions[i+2]]
			positions[positions[i+3]] = res
		case 99:
			return positions[0]
		default:
		}
	}
	return 0
}
