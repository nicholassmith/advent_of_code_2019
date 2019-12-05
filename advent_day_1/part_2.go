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
			mass, _ := strconv.Atoi(strings.TrimSpace(line))
			fuels = append(fuels, calculateFuelLoad(mass))
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		mass, _ := strconv.Atoi(strings.TrimSpace(line))

		fuels = append(fuels, calculateFuelLoad(mass))
	}
	fmt.Print(sumFuel(fuels))
}


func calculateFuelLoad(mass int) (int) {
	result := int(math.Floor(float64(mass / 3)) - 2)
	if (result > 0) {
		result = result + calculateFuelLoad(result)
		return int(result)
	}
	return 0
}

func sumFuel(fuels [] int) (int) {
	sum := 0 
	for _, fuel := range fuels {
		sum += fuel
	}
	return sum
}