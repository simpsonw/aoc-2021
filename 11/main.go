package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
		log.Fatalf("Usage: %s <filename> <steps>\n", os.Args[0])
	}
	filename := os.Args[1]
	steps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
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

	for i := 1; i <= steps; i++ {

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

		for row, v := range c.Octopuses {
			for col, octopus := range v {
				if octopus.Flashed {
					c.Flashes++
					c.Octopuses[row][col].Charge = 0
					c.Octopuses[row][col].Flashed = false
				}
			}
		}

		if i < 10 || i%10 == 0 {
			fmt.Printf("After step %d\n", i)
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
	}
	fmt.Printf("Counted %d flashes after %d steps\n", c.Flashes, steps)
}
