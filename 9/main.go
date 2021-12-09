package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("heightmap.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	heightMap := make([][]int, len(lines))
	numCols := len(lines[0])
	// Subtract the blank line at the end of the file
	numRows := len(lines) - 1
	var mins []int
	var riskLevelSum int
	for row, line := range lines {
		for _, v := range line {
			heightMap[row] = append(heightMap[row], int(v-'0'))
		}
	}

	for row, line := range heightMap {
		for col, v := range line {
			isMin := true

			//check left
			if col > 0 {
				isMin = isMin && v < line[col-1]
			}
			//check right
			if col < numCols-1 {
				isMin = isMin && v < line[col+1]
			}
			//check above
			if row > 0 {
				isMin = isMin && v < heightMap[row-1][col]
			}
			//check below
			if row < numRows-1 {
				isMin = isMin && v < heightMap[row+1][col]
			}
			if isMin {
				fmt.Printf("%d*", v)
				mins = append(mins, v)
				riskLevelSum += v + 1
			} else {
				fmt.Printf("%d ", v)
			}
		}
		fmt.Println("")
	}

	fmt.Printf("Sum of risk levels: %d\n", riskLevelSum)
}
