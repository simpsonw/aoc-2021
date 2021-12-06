package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("fish.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), ",")
	fish := make([]int, 9)
	for _, v := range lines {
		v := strings.TrimSpace(v)
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		fish[n]++
	}
	days := 256
	fmt.Println("Initial State:")
	for k, v := range fish {
		fmt.Printf("\tDay %d:Fish %d\n", k, v)
	}
	for i := 1; i <= days; i++ {
		newFish := make([]int, 9)
		var spawn int
		for k, v := range fish {
			if k != 0 {
				newFish[k-1] = v
			} else {
				spawn = v
			}
		}
		newFish[6] += spawn
		newFish[8] = spawn
		fish = newFish
	}

	total := 0
	for _, v := range fish {
		total += v
	}
	fmt.Printf("Total fish after %d days: %d\n", days, total)
}
