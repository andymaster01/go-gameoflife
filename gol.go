package main

import (
	"fmt"
)

func printBoard(board []int) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			p := (r * 5) + c
			fmt.Print(board[p])
		}
		fmt.Println("")
	}
}

func main() {
	var board = []int{
		0, 0, 0, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 0, 0, 0}

	printBoard(board)
}
