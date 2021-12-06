package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Orientation string

const (
	Vertical   Orientation = "vertical"
	Horizontal Orientation = "horizontal"
	Diagonal   Orientation = "diagonal"
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
	} else if l.Start.Y == l.End.Y {
		return Horizontal
	} else {
		return Diagonal
	}
}

func (l LineSegment) getPoints() []Coordinate {
	var points []Coordinate
	if l.getOrientation() == Vertical {
		to := l.End.Y
		from := l.Start.Y

		if l.Start.Y > l.End.Y {
			to = l.Start.Y
			from = l.End.Y
		}

		for i := from; i <= to; i++ {
			p := Coordinate{
				X: l.Start.X,
				Y: i,
			}
			points = append(points, p)
		}
	} else if l.getOrientation() == Horizontal {
		to := l.End.X
		from := l.Start.X

		if l.Start.X > l.End.X {
			to = l.Start.X
			from = l.End.X
		}

		for i := from; i <= to; i++ {
			p := Coordinate{
				X: i,
				Y: l.Start.Y,
			}
			points = append(points, p)
		}
	}
	return points
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(content), "\n")
	collisionPoints := make(map[Coordinate]bool, len(split))
	numCollisions := 0

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
		points := l.getPoints()
		for _, v := range points {
			repeat, ok := collisionPoints[v]
			if !ok {
				// This is the first time we've seen this point
				collisionPoints[v] = false
			} else if !repeat {
				// This is the second time we've seen this point.
				// Mark it as seen and increment collision count
				collisionPoints[v] = true
				numCollisions++
			}
		}
	}
	fmt.Printf("There are %d points that are intersected at least twice\n", numCollisions)
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
