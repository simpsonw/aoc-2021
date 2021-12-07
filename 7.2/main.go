package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("crabs.txt")
	lines := strings.Split(string(content), "\n")
	pos := strings.Split(lines[0], ",")
	crabs := make([]int, len(pos))
	for k, v := range pos {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		crabs[k] = n
	}
	n := sort.IntSlice(crabs)
	n.Sort()
	max := n[len(n)-1]
	bestFuelConsumption := -1
	bestPosition := -1
	for i := 0; i < max; i++ {
		fuelConsumption := 0
		for _, v := range crabs {
			delta := v - i
			if delta < 0 {
				delta *= -1
			}
			for j := 1; j <= delta; j++ {
				fuelConsumption += j
			}
		}
		if fuelConsumption < bestFuelConsumption || bestFuelConsumption == -1 {
			bestFuelConsumption = fuelConsumption
			bestPosition = i
		}
	}

	fmt.Printf("The best position is %d with a total fuel consumption of %d\n", bestPosition, bestFuelConsumption)
}
