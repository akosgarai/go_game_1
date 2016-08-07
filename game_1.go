package main

import (
	"fmt"
	"strings"
)

type Game struct {
	p1     string
	p2     string
	actual string
	board  [][]string
}

func main() {
	game := Game{"Player 1", "Player 2", "p1", initBoard(5, 7)}
	drawBoard(game.board)
	changeValue(game.board, 1, 3, "X")
	drawBoard(game.board)
}

func drawBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func changeValue(board [][]string, rowIndex, colIndex int, target string) [][]string {
	if board[rowIndex][colIndex] == "_" {
		board[rowIndex][colIndex] = target
	}
	return board
}

func initBoard(row, col int) [][]string {
	var board [][]string
	for j := 0; j < row; j++ {
		var row []string
		for i := 0; i < col; i++ {
			row = append(row, "_")
		}
		board = append(board, row)
	}
	return board
}
