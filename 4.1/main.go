package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input1.txt")
	split := strings.Split(string(content), "\n")

	var (
		allBoards            [][][]int
		currBoard            [][]int
		drawnNumbers         []int
		lastWinner           [][]int
		lastDrawnNumber      int
		lastDrawnNumberOrder int
	)
	numbers := split[0]
	for _, n := range strings.Split(numbers, ",") {
		num, _ := strconv.Atoi(n)
		drawnNumbers = append(drawnNumbers, num)
	}

	for k, v := range split {
		if k < 2 {
			//first two lines contain winning numbers and blank line
			continue
		}
		if v == "" {
			//Boards are separated by blank lines
			allBoards = append(allBoards, currBoard)
			currBoard = make([][]int, 0)
			continue
		}

		nums := strings.Fields(v)
		row := make([]int, 0)
		for _, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, i)
		}
		currBoard = append(currBoard, row)
	}

	for _, board := range allBoards {
		fmt.Printf("Board: %v\n", board)
		for k, number := range drawnNumbers {
			markSquares(board, number)
			if isWinner(board) {
				if k > lastDrawnNumberOrder {
					lastDrawnNumber = number
					lastWinner = board
					lastDrawnNumberOrder = k
				}
				break
			}
		}
	}
	score := getScore(lastWinner, lastDrawnNumber)
	fmt.Printf("Last Winner: %v\nLast Score: %d\nOrder:%d\n", lastWinner, score, lastDrawnNumberOrder)
}

func getScore(board [][]int, multiplier int) int {
	score := 0
	for _, row := range board {
		for _, v := range row {
			if v != -1 {
				score += v
			}
		}
	}
	return score * multiplier
}

func markSquares(board [][]int, number int) {
	for _, row := range board {
		for k, v := range row {
			if v == number {
				row[k] = -1
			}
		}
	}
}

func isWinner(board [][]int) bool {
	for _, v := range board {
		rowWin := true
		for _, j := range v {
			if j != -1 {
				rowWin = false
			}
		}
		if rowWin {
			return true
		}
	}

	for col := 0; col < len(board[0]); col++ {
		colWin := true
		for i := 0; i < len(board[0]); i++ {
			if board[i][col] != -1 {
				colWin = false
			}
		}
		if colWin {
			return true
		}
	}
	return false
}
