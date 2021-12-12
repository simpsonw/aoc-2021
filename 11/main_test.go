package main

import (
	"testing"
)

func assertOctopusPointerValue(t testing.TB, got *Octopus, want Octopus) {
	t.Helper()
	if *got != want {
		t.Fatalf("got %d wanted %d\n", *got, want)
	}
}

func TestNeighbors(t *testing.T) {
	ints := [][]int{
		{1, 1, 1, 1, 1},
		{1, 9, 9, 9, 1},
		{1, 9, 1, 9, 1},
		{1, 9, 9, 9, 1},
		{1, 1, 1, 1, 1},
	}
	o := make([][]Octopus, len(ints))
	for row, v := range ints {
		o[row] = make([]Octopus, len(v))
		for col, i := range v {
			o[row][col] = Octopus{
				Row:    row,
				Col:    col,
				Charge: i,
			}
		}
	}

	c := Cave{
		Rows:      len(o),
		Cols:      len(o[0]),
		Octopuses: o,
	}

	t.Run("Top left", func(t *testing.T) {
		got := c.GetNeighbors(0, 0)
		want := []Octopus{
			{Row: 0, Col: 1, Charge: 1},
			{Row: 1, Col: 1, Charge: 9},
			{Row: 1, Col: 0, Charge: 1},
		}
		for k, v := range want {
			assertOctopusPointerValue(t, got[k], v)
		}
	})

	t.Run("Bottom left", func(t *testing.T) {
		got := c.GetNeighbors(4, 0)
		want := []Octopus{
			{Row: 4, Col: 1, Charge: 1},
			{Row: 3, Col: 1, Charge: 9},
			{Row: 3, Col: 0, Charge: 1},
		}
		for k, v := range want {
			assertOctopusPointerValue(t, got[k], v)
		}
	})

	t.Run("Top right", func(t *testing.T) {
		got := c.GetNeighbors(0, 4)
		want := []Octopus{
			{Row: 0, Col: 3, Charge: 1},
			{Row: 1, Col: 3, Charge: 9},
			{Row: 1, Col: 4, Charge: 1},
		}
		for k, v := range want {
			assertOctopusPointerValue(t, got[k], v)
		}
	})

	t.Run("Bottom right", func(t *testing.T) {
		got := c.GetNeighbors(4, 4)
		want := []Octopus{
			{Row: 4, Col: 3, Charge: 1},
			{Row: 3, Col: 3, Charge: 9},
			{Row: 3, Col: 4, Charge: 1},
		}
		for k, v := range want {
			assertOctopusPointerValue(t, got[k], v)
		}
	})

	t.Run("Center", func(t *testing.T) {
		got := c.GetNeighbors(2, 2)
		want := []Octopus{
			{Row: 2, Col: 1, Charge: 9},
			{Row: 1, Col: 1, Charge: 9},
			{Row: 3, Col: 1, Charge: 9},
			{Row: 2, Col: 3, Charge: 9},
			{Row: 1, Col: 3, Charge: 9},
			{Row: 3, Col: 3, Charge: 9},
			{Row: 1, Col: 2, Charge: 9},
			{Row: 3, Col: 2, Charge: 9},
		}
		for k, v := range want {
			assertOctopusPointerValue(t, got[k], v)
		}
	})
}
