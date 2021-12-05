package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("depths.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var depths []int

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, curr)
	}

	var increases, a, b int
	numDepths := len(depths)
	for k, _ := range depths {
		if k+3 >= numDepths {
			break
		}
		for i := k; i <= k+3; i++ {
			if i == k {
				a += depths[i]
			} else if i < k+3 {
				a += depths[i]
				b += depths[i]
			} else {
				b += depths[i]
			}
		}
		if b > a {
			fmt.Printf("%d is greater than %d\n", b, a)
			increases++
		}
		b, a = 0, 0
	}

	fmt.Printf("Total increases: %d\n", increases)
}
