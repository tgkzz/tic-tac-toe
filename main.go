package main

import "fmt"

const (
	N     = 3
	Empty = 0
)

type Board [N][N]rune

func NewGame(b *Board) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			b[i][j] = '-'
		}
	}
}

func DrawBoard(b *Board) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(string(b[i][j]))
		}
		fmt.Println()
	}
}

func playerWon(b *Board) rune {
	for i := 0; i < N; i++ {
		if b[i][0] == b[i][1] && b[i][1] == b[i][2] && b[i][0] != '-' {
			return b[i][0]
		}
	}

	for j := 0; j < N; j++ {
		if b[0][j] == b[1][j] && b[1][j] == b[2][j] && b[0][j] != '-' {
			return b[0][j]
		}
	}

	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[0][0] != '-' {
		return b[0][0]
	}
	if b[2][0] == b[1][1] && b[1][1] == b[0][2] && b[2][0] != '-' {
		return b[2][0]
	}

	return ' '
}

func boardIsFull(b *Board) bool {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if b[i][j] == '-' {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println("welcome to my tic-tac-toe")

	var b Board

	NewGame(&b)

	var player1 string
	var player2 string

	fmt.Println("Player 1, enter your name")
	fmt.Scanln(&player1)
	fmt.Println("Player 2, enter your name")
	fmt.Scanln(&player2)

	moveP1 := true
	gameOver := false

	for !gameOver {
		DrawBoard(&b)
		if moveP1 {
			fmt.Println(player1 + "'s turn")
		} else {
			fmt.Println(player2 + "'s turn")
		}

		var ch rune

		if moveP1 {
			ch = 'X'
		} else {
			ch = 'O'
		}

		var row int
		var col int

		for true {
			fmt.Println("Enter a row number (0, 1, or 2): ")
			fmt.Scanln(&row)
			fmt.Println("Enter a col number (0, 1, or 2): ")
			fmt.Scanln(&col)

			if row < 0 || col < 0 || row > 2 || col > 2 {
				fmt.Println("This position is out of the bounds!")
			} else if b[row][col] != '-' {
				fmt.Println("This position is already taken!")
			} else {
				break
			}
		}

		b[row][col] = ch

		if playerWon(&b) == 'X' {
			fmt.Println(player1 + " has won")
			gameOver = true
		} else if playerWon(&b) == 'O' {
			fmt.Println(player2 + " has won")
			gameOver = true
		} else {
			if boardIsFull(&b) {
				fmt.Println("it is draw")
				gameOver = true
			} else {
				moveP1 = !moveP1
			}
		}
	}

	DrawBoard(&b)
}
