package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")
	ySegments := make(map[int]Coordinate, len(split))
	xSegments := make(map[int]Coordinate, len(split))

	for _, v := range split {
		vals := strings.Fields(v)
		if len(vals) != 3 {
			break
		}
		x1, y1, err := getCoordinatePair(vals[0])
		if err != nil {
			log.Fatal(err)
		}
		coordinate1 := &Coordinate{
			X: x1,
			Y: y1,
		}
		xSegments[x1] = *coordinate1
		ySegments[y1] = *coordinate1
	}
}

func getCoordinatePair(s string) (int, int, error) {
	split := strings.Split(s, ",")
	x, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}
