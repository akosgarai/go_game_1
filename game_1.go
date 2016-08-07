package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Game struct {
	p1    string
	p2    string
	state string
	board [][]string
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func userInputForm(question string) string {
	clearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question, "\n")
	text, _ := reader.ReadString('\n')
	return text
}
func initGame() Game {
	var p1Name, p2Name string
	var boardRow, boardCol int
	p1Name = userInputForm("First player's name?")
	p2Name = userInputForm("Second player's name?")
	boardRow = 5
	boardCol = 5
	return Game{p1Name, p2Name, "p1", initBoard(boardRow, boardCol)}
}

func main() {
	game := initGame()
	clearScreen()
	drawBoard(game.board)
	changeValue(game.board, 1, 3, "X")
	clearScreen()
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
