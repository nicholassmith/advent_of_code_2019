package main

import(
	"bufio"
	"os"
	"fmt"
	"log"
	"io"
	"strconv"
	"math"
	"strings"
)

func main() {
	var fuels[] int

	f, err := os.Open("mass.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			fuels = append(fuels, calculateFuelLoad(line))
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fuels = append(fuels, calculateFuelLoad(line))
	}
	fmt.Print(sumFuel(fuels))
}


func calculateFuelLoad(mass string) (int) {
	numMass, err := strconv.Atoi(strings.TrimSpace(mass))
	if err != nil {
		log.Fatal(err)
	}
	result := math.Floor(float64(numMass / 3)) - 2
	return int(result)
}

func sumFuel(fuels [] int) (int) {
	sum := 0 
	for _, fuel := range fuels {
		sum += fuel
	}
	return sum
}