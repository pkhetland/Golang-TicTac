package tictac

import (
	"bytes"
	"fmt"
)

const (
	dimension = 3
	Nobody    = dimension + 1
	Player1   = 0
	Player2   = 1
)

type Player int

func (p Player) String() string {
	switch p {
	case Player1:
		return "X"
	case Player2:
		return "O"
	default:
		return " "
	}
}

type Game struct {
	board         []Player
	currentPlayer Player
	moveCount     int
	gameOver      bool
	hasWinner     bool
}

func (g Game) String() string {
	var result bytes.Buffer
	for i, player := range g.board {
		result.WriteString(fmt.Sprintf(" %s ", player))
		if (i+1)%dimension != 0 {
			result.WriteString("|")
			continue
		}

		result.WriteString("\n")
		if i != len(g.board)-1 {
			for j := 1; j < dimension+1; j++ {
				result.WriteString("---")

			}
			result.WriteString("\n")
		}
	}
	return result.String()
}

// NewGame is a factory function for a new game of Tic-Tac-Toe.
func NewGame() *Game {
	b := make([]Player, dimension*dimension)
	for i := range b {
		b[i] = Nobody
	}
	return &Game{board: b}
}

func (g *Game) Play(x, y int) error {
	switch {
	case g.gameOver:
		return fmt.Errorf("Game already over")
	case x < 0 || x > dimension-1 || y < 0 || y > dimension-1:
		return fmt.Errorf("Invalid coordinates")
	case g.board[index(x, y)] != Nobody:
		return fmt.Errorf("Field already marked")
	}

	g.board[index(x, y)] = g.currentPlayer
	g.moveCount++

	g.checkStatus(x, y)
	if !g.gameOver {
		g.currentPlayer = (g.currentPlayer + 1) % 2
	}
	return nil
}

func (g *Game) checkStatus(x, y int) {
	switch {
	case g.moveCount < dimension+2:
		return
	case g.moveCount == dimension*dimension:
		g.gameOver = true
	}

	winCase := dimension * g.currentPlayer

	var row, col, dia, rdia Player
	for i := 0; i < dimension; i++ {
		row += g.board[index(i, y)]
		col += g.board[index(x, i)]
		dia += g.board[index(i, i)]
		rdia += g.board[index(dimension-(i+1), i)]
	}

	if checkFor(winCase, row, col, dia, rdia) {
		g.gameOver = true
		g.hasWinner = true
	}
}

func checkFor(value Player, items ...Player) bool {
	for _, item := range items {
		if value == item {
			return true
		}
	}
	return false
}

func (g Game) FieldValue(x, y int) Player {
	return g.board[index(x, y)]
}

func index(x, y int) int {
	return x + dimension*y
}
