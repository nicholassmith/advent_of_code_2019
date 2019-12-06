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

	part1pos := make([]int, len(positions))
	part2pos := make([]int, len(positions))

	copy(part1pos, positions)
	copy(part2pos, positions)

	fmt.Printf("part 1: %d\n", compute(part1pos))
	fmt.Printf("part 2: %d\n", part2(part2pos))
}

func part2(positions []int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			clone := make([]int, len(positions))
			copy(clone, positions)
			clone[1] = noun
			clone[2] = verb
			if compute(clone) == 19690720 {
				return (100 * noun) + verb
			}
		}
	}

	return 0
}

func compute(positions []int) int {
	for i := 0; i < len(positions); i += 4 {
		opcode := positions[i]
		pos1 := positions[i+1]
		pos2 := positions[i+2]
		store := positions[i+3]
		switch opcode {
		case 1:
			positions[store] = positions[pos1] + positions[pos2]
		case 2:
			positions[store] = positions[pos1] * positions[pos2]
		case 99:
			return positions[0]
		default:
		}
	}
	return 0
}
