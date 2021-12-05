package main

import (
	"errors"
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
