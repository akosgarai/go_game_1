package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const timeDefault int = 120
const targetNum int = 5
const numOfRows int = 10
const numOfCols int = 10

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
	row  int
	col  int
	time int
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func getStringUserInputAfterPrintedText(textToPrint string) string {
	clearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(textToPrint, "\n")
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
func userRestartHandler(u User) bool {
	var response string
	var statusmsg string
	statusmsg = "Player " + u.name + " won\nNew game?"
	response = getStringUserInputAfterPrintedText(statusmsg)
	if response == "no" {
		return false
	}
	return true
}
func initGame() Game {
	var p1Name, p2Name string
	p1Name = ""
	for p1Name == "" {
		p1Name = getStringUserInputAfterPrintedText("First player's name?")
	}
	p2Name = ""
	for p2Name == "" {
		p2Name = getStringUserInputAfterPrintedText("Second player's name?")
	}
	return Game{User{p1Name, timeDefault, 0, Target{-1, -1, -1}}, User{p2Name, timeDefault, 0, Target{-1, -1, -1}}, "p1", initBoard(numOfRows, numOfCols)}
}

func drawUserInfo(p1, p2 User, state string) {
	switch state {
	case "p1":
		fmt.Print(p1.name, " turn.\n")
	case "p2":
		fmt.Print(p2.name, " turn.\n")
	case "end":
		var winner string
		if p1.points > p2.points {
			winner = p1.name
		} else if p1.points < p2.points {
			winner = p2.name
		} else {
			winner = "Points are equal, both players are looser."
		}
		fmt.Print("The winner is: ", winner, "\n")
	}
}

func getElapsedTime(t1, t2 int64) int {
	diff := t2 - t1
	diffInSec := diff / int64(time.Second)
	return int(diffInSec)
}

func getNumericUserInputWithPrintedText(userText string) int {
	var i int
	fmt.Print(userText, " ")
	_, err := fmt.Scanf("%d", &i)
	if err == nil {
		return i
	} else {
		return -1
	}
}

func userInputHandler() Target {
	var i, j int
	t := time.Now().UTC().UnixNano()
	i = getNumericUserInputWithPrintedText("Row:")
	if i > -1 {
		j = getNumericUserInputWithPrintedText("Col:")
	} else {
		j = -1
	}
	t2 := time.Now().UTC().UnixNano()
	return Target{i, j, getElapsedTime(t, t2)}
}

func restartGame(game Game) {
	game.p1.timeLeft = timeDefault
	game.p2.timeLeft = timeDefault
	game.board = initBoard(numOfRows, numOfCols)
}

func drawScreen(game Game) {
	clearScreen()
	drawMenu(game.p1, game.p2)
	drawBoard(game.board)
	drawUserInfo(game.p1, game.p2, game.state)
}

func game_fsm(game Game) {
	drawScreen(game)
	for game.state != "end" {
		switch game.state {
		case "p1":
			t := userInputHandler()
			game.p1.timeLeft = game.p1.timeLeft - t.time
			if game.p1.timeLeft < 0 {
				game.state = "p1_fall"
				continue
			}
			if t.row > -1 && t.col > -1 {
				if !changeValue(game.board, t.row, t.col, "X") {
					game.state = "p1"
				} else if isWinnerStep(t, game.board) {
					game.state = "p1_win"
				} else {
					game.state = "p2"
				}
			}
		case "p2":
			t := userInputHandler()
			game.p2.timeLeft = game.p2.timeLeft - t.time
			if game.p2.timeLeft < 0 {
				game.state = "p2_fall"
				continue
			}
			if t.row > -1 && t.col > -1 {
				if !changeValue(game.board, t.row, t.col, "O") {
					game.state = "p2"
				} else if isWinnerStep(t, game.board) {
					game.state = "p2_win"
				} else {
					game.state = "p1"
				}
			}
		case "p2_fall":
			newGame := userRestartHandler(game.p1)
			if newGame {
				game.p1.points = game.p1.points + 1
				restartGame(game)
				game.state = "p2"
			} else {
				game.state = "end"
			}
		case "p1_fall":
			newGame := userRestartHandler(game.p2)
			if newGame {
				game.p2.points = game.p2.points + 1
				restartGame(game)
				game.state = "p1"
			} else {
				game.state = "end"
			}
		case "p1_win":
			game.p1.points = game.p1.points + 1
			restartGame(game)
			newGame := userRestartHandler(game.p1)
			if newGame {
				game.state = "p1"
			} else {
				game.state = "end"
			}

		case "p2_win":
			game.p2.points = game.p2.points + 1
			restartGame(game)
			newGame := userRestartHandler(game.p2)
			if newGame {
				game.state = "p2"
			} else {
				game.state = "end"
			}
		}
		drawScreen(game)
	}
}

func isWinnerStep(t Target, b [][]string) bool {
	var row, col int = t.row, t.col
	var expected string = b[row][col]
	var j, k int = 1, 1
	var totalHit int = 1
	direction := 1 //1 - up-down, 2 - left-right, 3 - up,right-down,left, 4 - up,left-down,right
	for direction < 5 {
		switch direction {
		case 1:
			totalHit = 1
			j = 1
			for row-j >= 0 && b[row-j][col] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				j = j + 1
			}
			k = 1
			for row+k < numOfRows && b[row+k][col] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				k = k + 1
			}
			direction = 2
		case 2:
			totalHit = 1
			j = 1
			for col-j >= 0 && b[row][col-j] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				j = j + 1
			}
			k = 1
			for col+k < numOfCols && b[row][col+k] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				k = k + 1
			}
			direction = 3
		case 3:
			totalHit = 1
			j = 1
			for row+j < numOfRows && col+j < numOfCols && b[row+j][col+j] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				j = j + 1
			}
			k = 1
			for row-k >= 0 && col-k >= 0 && b[row-k][col-k] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				k = k + 1
			}
			direction = 4
		case 4:
			totalHit = 1
			j = 1
			for row-j >= 0 && col+j < numOfRows && b[row-j][col+j] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				j = j + 1
			}
			k = 1
			for row+k < numOfRows && col-k >= 0 && b[row+k][col-k] == expected {
				totalHit = totalHit + 1
				if totalHit >= targetNum {
					return true
				}
				k = k + 1
			}
			direction = 5
			return false
		}
	}
	return false
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

func changeValue(board [][]string, rowIndex, colIndex int, target string) bool {
	if board[rowIndex][colIndex] == "_" {
		board[rowIndex][colIndex] = target
		return true
	}
	return false
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
