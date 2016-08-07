package main

import (
    "fmt"
    "strings"
)

func main () {
        board := initBoard(5,7)
        drawBoard(board)
}

func drawBoard (board [][]string) {
        for i := 0; i < len(board); i++ {
                fmt.Printf("%s\n", strings.Join(board[i], " "))
        }
}

func initBoard (row, col int) [][]string {
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

