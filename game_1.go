package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Game struct {
	p1    User
	p2    User
	state string
	board [][]string
}

type User struct {
	name     string
	timeLeft int
	points   int
	t        Target
}

type Target struct {
	row int
	col int
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
	return strings.Replace(text, "\n", "", -1)
}
func initGame() Game {
	var p1Name, p2Name string
	var boardRow, boardCol int
	p1Name = userInputForm("First player's name?")
	p2Name = userInputForm("Second player's name?")
	boardRow = 10
	boardCol = 10
	return Game{User{p1Name, 120, 0, Target{-1, -1}}, User{p2Name, 120, 0, Target{-1, -1}}, "p1", initBoard(boardRow, boardCol)}
}

func drawUserInfo(game Game) {
	//var i int
	//var q []string
	switch game.state {
	case "p1":
		fmt.Print(game.p1.name, " turn.")
	case "p2":
		fmt.Print(game.p2.name, " turn.")
	case "end":
		fmt.Print("end\n")
	}
}

func userInputHandler(game Game) Target {
	var i, j int
	fmt.Print("Row: ")
	_, err := fmt.Scanf("%d", &i)
	if err == nil {
		fmt.Print("Col: ")
		_, err2 := fmt.Scanf("%d", &j)

		if err2 == nil {
			if game.board[i][j] == "_" {
				return Target{i, j}
			}
		}
	}
	return Target{-1, -1}
}

func game_fsm(game Game) {
	clearScreen()
	drawMenu(game.p1, game.p2)
	drawBoard(game.board)
	drawUserInfo(game)
	for game.state != "end" {
		switch game.state {
		case "p1":
			t := userInputHandler(game)
			if t.row > -1 && t.col > -1 {
				changeValue(game.board, t.row, t.col, "X")
				game.state = "p2"
			}
		case "p2":
			t := userInputHandler(game)
			if t.row > -1 && t.col > -1 {
				changeValue(game.board, t.row, t.col, "O")
				game.state = "p1"
			}
		}
		clearScreen()
		drawMenu(game.p1, game.p2)
		drawBoard(game.board)
		drawUserInfo(game)
	}
}

func main() {
	game := initGame()
	game_fsm(game)
}

func drawMenu(u1, u2 User) {
	fmt.Print(u1.name, " ", u1.timeLeft, " ", u1.points, " ", u2.points, " ", u2.timeLeft, " ", u2.name, "\n")
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
