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

	firstLine := true
	var prev, curr, increases int
	for scanner.Scan() {
		if firstLine {
			prev, err = strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			firstLine = false
			continue
		}

		curr, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Is %d greater than %d?\n", curr, prev)
		if curr > prev {
			increases++
			fmt.Printf("\tYes!  increases: %d\n", increases)
		}
		prev = curr
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There were %d increases\n", increases)
}
