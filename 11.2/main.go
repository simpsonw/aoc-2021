package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Octopus struct {
	Row     int
	Col     int
	Charge  int
	Flashed bool
}

func (o *Octopus) String() string {
	return fmt.Sprintf("Row: %d Col: %d Charge: %d\n", o.Row, o.Col, o.Charge)
}

type Cave struct {
	Rows      int
	Cols      int
	Octopuses [][]Octopus
	Flashes   int
}

func (c *Cave) Flash(row, col int) {
	o := &c.Octopuses[row][col]
	if o.Flashed == true {
		return
	}
	o.Flashed = true
	neighbors := c.GetNeighbors(row, col)
	for len(neighbors) > 0 {
		o = neighbors[0]
		o.Charge++
		if o.Charge > 9 && !o.Flashed {
			o.Flashed = true
			neighbors = append(neighbors, c.GetNeighbors(o.Row, o.Col)...)
		}
		neighbors = neighbors[1:]
	}
}

func (c *Cave) GetNeighbors(row, col int) []*Octopus {
	var neighbors []*Octopus
	// Left
	if col > 0 {
		neighbors = append(neighbors, &c.Octopuses[row][col-1])

		// Top Left
		if row > 0 {
			neighbors = append(neighbors, &c.Octopuses[row-1][col-1])
		}
		// Bottom Left
		if row < c.Rows-1 {
			neighbors = append(neighbors, &c.Octopuses[row+1][col-1])
		}
	}
	// Right
	if col < c.Cols-1 {
		neighbors = append(neighbors, &c.Octopuses[row][col+1])

		// Top Right
		if row > 0 {
			neighbors = append(neighbors, &c.Octopuses[row-1][col+1])
		}
		// Bottom Right
		if row < c.Rows-1 {
			neighbors = append(neighbors, &c.Octopuses[row+1][col+1])
		}
	}
	// Top
	if row > 0 {
		neighbors = append(neighbors, &c.Octopuses[row-1][col])
	}
	// Bottom
	if row < c.Rows-1 {
		neighbors = append(neighbors, &c.Octopuses[row+1][col])
	}
	return neighbors
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>\n", os.Args[0])
	}
	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	o := make([][]Octopus, len(lines)-1)
	for row, l := range lines {
		if l == "" {
			continue
		}
		o[row] = make([]Octopus, len(l))
		for col, v := range l {
			o[row][col] = Octopus{
				Row:    row,
				Col:    col,
				Charge: int(v - '0'),
			}
		}
	}
	c := Cave{
		Rows:      len(o),
		Cols:      len(o[0]),
		Octopuses: o,
	}

	fmt.Println("Before any steps")
	for _, row := range c.Octopuses {
		for _, octopus := range row {
			if octopus.Charge > 0 {
				fmt.Printf(" %d ", octopus.Charge)
			} else {
				fmt.Printf("[%d]", octopus.Charge)
			}
		}
		fmt.Println("")
	}
	fmt.Println("")

	step := 1
	for {

		for row, v := range c.Octopuses {
			for col, _ := range v {
				c.Octopuses[row][col].Charge++
			}
		}

		for row, v := range c.Octopuses {
			for col, octopus := range v {
				if octopus.Charge > 9 {
					c.Flash(row, col)
				}
			}
		}

		allFlashed := true
		for row, v := range c.Octopuses {
			for col, octopus := range v {
				allFlashed = allFlashed && c.Octopuses[row][col].Flashed
				if octopus.Flashed {
					c.Octopuses[row][col].Charge = 0
					c.Octopuses[row][col].Flashed = false
				}
			}
		}
		if allFlashed {
			fmt.Printf("All the octopuses flashed in step %d\n", step)
			break
		}

		if step < 10 || step%10 == 0 {
			fmt.Printf("After step %d\n", step)
			for _, row := range c.Octopuses {
				for _, octopus := range row {
					if octopus.Charge > 0 {
						fmt.Printf(" %d ", octopus.Charge)
					} else {
						fmt.Printf("[%d]", octopus.Charge)
					}
				}
				fmt.Println("")
			}
			fmt.Println("")

		}
		step++
	}
}
