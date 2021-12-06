package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Orientation string

const (
	Vertical   Orientation = "vertical"
	Horizontal Orientation = "horizontal"
)

type Coordinate struct {
	X int
	Y int
}

type LineSegment struct {
	Start Coordinate
	End   Coordinate
}

func (l LineSegment) getOrientation() Orientation {
	if l.Start.X == l.End.X {
		return Vertical
	} else {
		return Horizontal
	}
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")
	lines := make([]LineSegment, len(split))

	for _, v := range split {
		vals := strings.Fields(v)
		if len(vals) != 3 {
			break
		}
		x1, y1, err := getCoordinatePair(vals[0])
		if err != nil {
			log.Fatal(err)
		}

		x2, y2, err := getCoordinatePair(vals[2])
		if err != nil {
			log.Fatal(err)
		}

		l := LineSegment{
			Start: Coordinate{
				X: x1,
				Y: y1,
			},
			End: Coordinate{
				X: x2,
				Y: y2,
			},
		}
		lines = append(lines, l)
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
