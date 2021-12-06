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
	fish := make([]int, len(lines))
	for k, v := range lines {
		v := strings.TrimSpace(v)
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		fish[k] = n
	}
	days := 80
	fmt.Printf("Initial state: %v\n", fish)
	for i := 1; i <= days; i++ {
		for k, v := range fish {
			if v == 0 {
				fish[k] = 6
				fish = append(fish, 8)
			} else {
				fish[k]--
			}
		}
		fmt.Printf("After\t%d day(s):\t%d\n", i, len(fish))
	}
	fmt.Printf("Total fish after %d days: %d\n", days, len(fish))
}
