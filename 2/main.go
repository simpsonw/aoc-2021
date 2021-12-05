package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var horizontalPosition, depth int
	f, err := os.Open("moves.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		if len(s) != 2 {
			log.Fatalf("Invalid line: %s\n", line)
		}
		value, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		switch s[0] {
		case "forward":
			horizontalPosition += value
			break
		case "down":
			depth += value
			break
		case "up":
			depth -= value
			break
		default:
			log.Fatalf("Unrecognized option: %s\n", s[0])
		}
	}
	fmt.Printf("Horizontal Position: %d\nDepth: %d\n", horizontalPosition, depth)

}
