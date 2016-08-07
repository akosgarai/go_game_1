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
	boardRow = 10
	boardCol = 10
	return Game{p1Name, p2Name, "p1", initBoard(boardRow, boardCol)}
}

func getUserInputs() {
	var i int
	_, err := fmt.Scanf("%d", &i)
}

func game_fsm(game Game) {
	clearScreen()
	drawBoard(game.board)
	for game.state != "end" {
		switch game.state {
		case "p1":
			changeValue(game.board, 1, 3, "X")
			game.state = "p2"
		case "p2":
			changeValue(game.board, 1, 4, "O")
			game.state = "end"
		}
		clearScreen()
		drawBoard(game.board)
	}
}

func main() {
	game := initGame()
	game_fsm(game)
}

func drawBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s | %d\n", strings.Join(board[i], " "), i)
	}
	for j := 0; j < len(board[0]); j++ {
		fmt.Print(j, " ")
	}
	fmt.Print("\n")
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
