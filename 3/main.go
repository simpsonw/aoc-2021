package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	bitSize := len(lines[0])
	var nums []int64
	for _, v := range lines {
		b, err := strconv.ParseInt(v, 2, 64)
		if err == nil {
			nums = append(nums, b)
		}
	}
	co2nums := nums

	for i := 1; i <= bitSize; i++ {
		var set, unset []int64
		if len(nums) == 1 {
			break
		}
		for _, v := range nums {
			if v&(1<<(bitSize-i)) > 0 {
				set = append(set, v)
			} else {
				unset = append(unset, v)
			}
		}
		if len(set) > len(unset) {
			nums = set
		} else if len(set) < len(unset) {
			nums = unset
		} else {
			nums = set
		}
	}
	fmt.Printf("o2: %05b\n", nums[0])

	for i := 1; i <= bitSize; i++ {
		var set, unset []int64
		if len(co2nums) == 1 {
			break
		}
		for _, v := range co2nums {
			if v&(1<<(bitSize-i)) > 0 {
				set = append(set, v)
			} else {
				unset = append(unset, v)
			}
		}
		if len(set) < len(unset) {
			co2nums = set
		} else if len(set) > len(unset) {
			co2nums = unset
		} else {
			co2nums = unset
		}
	}
	fmt.Printf("co2: %05b\n", co2nums[0])
	fmt.Printf("Life support rating: %d\n", nums[0]*co2nums[0])
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
