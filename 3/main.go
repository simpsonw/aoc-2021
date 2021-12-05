package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var setBits []uint
	var allLines []string
	var numLines, gammaRate, epsilonRate uint
	for scanner.Scan() {
		line := scanner.Text()
		allLines = append(allLines, line)
		if err != nil {
			log.Fatal(err)
		}

		if setBits == nil {
			setBits = make([]uint, len(line))
		}

		for i, c := range line {
			if c == '1' {
				setBits[i]++
			}
		}
		numLines++
	}
	offset := len(setBits) - 1
	for k, v := range setBits {
		if v >= numLines/2 {
			gammaRate |= (1 << (offset - k))
		} else {
			epsilonRate |= (1 << (offset - k))
		}
	}

	o2prefix := "10111"

	sort.Sort(sort.StringSlice(allLines))
	for _, v := range allLines {
		fmt.Println(v)
	}
	fmt.Printf("o2prefix: %s\n", o2prefix)
	o2 := splitSliceByPrefixHelper(0, o2prefix, allLines)
	fmt.Printf("The Oxygen Level is: %s\n", o2)
}

func splitSliceByPrefixHelper(index int, prefix string, slice []string) string {
	fmt.Printf("full prefix: %s\t prefix segment: %c index: %d\n", prefix, prefix[index], index)
	if len(prefix)-1 == index || len(slice) == 1 {
		return slice[0]
	}
	for k, v := range slice {
		if prefix[index] == v[index] {
			fmt.Printf("\tNew slice %s\n", slice[k:])
			return splitSliceByPrefixHelper(index+1, prefix, slice[k:])
		}
	}
	return slice[0]
}

func main_part1() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var setBits []uint
	var numLines, gammaRate, epsilonRate uint
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		if setBits == nil {
			setBits = make([]uint, len(line))
		}

		fmt.Printf("%s\n", line)
		for i, c := range line {
			if c == '1' {
				setBits[i]++
			}
		}
		numLines++
	}
	offset := len(setBits) - 1
	for k, v := range setBits {
		if v >= numLines/2 {
			gammaRate |= (1 << (offset - k))
		} else {
			epsilonRate |= (1 << (offset - k))
		}
	}

	fmt.Printf("gamma rate: %d (%05b)\n", gammaRate, gammaRate)
	fmt.Printf("epsilon rate: %d (%05b)\n", epsilonRate, epsilonRate)
	fmt.Printf("Power consumption: %d\n", gammaRate*epsilonRate)
}
