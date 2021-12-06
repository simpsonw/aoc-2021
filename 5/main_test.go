package main

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if !errors.Is(got, want) {
		t.Fatalf("got %q, want %q\n", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("got an error but didn't want one: %s\n", got)
	}
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got %d wanted %d\n", got, want)
	}
}

func TestGetCoordinatePairs(t *testing.T) {
	t.Run("Valid coordinate pair", func(t *testing.T) {
		input := "3,4"
		expectedX := 3
		expectedY := 4

		x, y, err := getCoordinatePair(input)

		assertNoError(t, err)
		assertInt(t, x, expectedX)
		assertInt(t, y, expectedY)
	})

	t.Run("Non-numeric coordinate pair", func(t *testing.T) {
		input := "a,b"
		expectedX := 0
		expectedY := 0
		x, y, err := getCoordinatePair(input)

		assertError(t, err, strconv.ErrSyntax)
		assertInt(t, x, expectedX)
		assertInt(t, y, expectedY)
	})
}

func TestLineOrientation(t *testing.T) {
	t.Run("Horizontal", func(t *testing.T) {
		l := LineSegment{
			Start: Coordinate{
				X: 1,
				Y: 10,
			},
			End: Coordinate{
				X: 10,
				Y: 10,
			},
		}
		want := l.getOrientation()
		got := Horizontal
		if got != want {
			t.Errorf("got %s wanted %s", got, want)
		}
	})
	t.Run("Vertical", func(t *testing.T) {
		l := LineSegment{
			Start: Coordinate{
				X: 1,
				Y: 1,
			},
			End: Coordinate{
				X: 1,
				Y: 10,
			},
		}
		want := l.getOrientation()
		got := Vertical
		if got != want {
			t.Errorf("got %s wanted %s", got, want)
		}
	})
	// Single point lines default to being vertical
	t.Run("Single point", func(t *testing.T) {
		l := LineSegment{
			Start: Coordinate{
				X: 1,
				Y: 1,
			},
			End: Coordinate{
				X: 1,
				Y: 1,
			},
		}
		want := l.getOrientation()
		got := Vertical
		if got != want {
			t.Errorf("got %s wanted %s", got, want)
		}
	})
}

func TestGetPoints(t *testing.T) {

	t.Run("Horizontal Line", func(t *testing.T) {
		l := LineSegment{
			Start: Coordinate{
				X: 1,
				Y: 1,
			},
			End: Coordinate{
				X: 1,
				Y: 3,
			},
		}
		got := l.getPoints()
		want := []Coordinate{
			{X: 1, Y: 1},
			{X: 1, Y: 2},
			{X: 1, Y: 3},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v\n", got, want)
		}
	})

	t.Run("Diagonal Line", func(t *testing.T) {
		l := LineSegment{
			Start: Coordinate{
				X: 1,
				Y: 1,
			},
			End: Coordinate{
				X: 3,
				Y: 3,
			},
		}
		got := l.getPoints()
		want := []Coordinate{
			{X: 1, Y: 1},
			{X: 2, Y: 2},
			{X: 3, Y: 3},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v\n", got, want)
		}
	})

	t.Run("Diagonal Line Down", func(t *testing.T) {
		l := LineSegment{
			Start: Coordinate{
				X: 9,
				Y: 7,
			},
			End: Coordinate{
				X: 7,
				Y: 9,
			},
		}
		got := l.getPoints()
		want := []Coordinate{
			{X: 9, Y: 7},
			{X: 8, Y: 8},
			{X: 7, Y: 9},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v\n", got, want)
		}
	})
}
