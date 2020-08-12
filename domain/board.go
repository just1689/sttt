package domain

import (
	"errors"
	"fmt"
	"github.com/just1689/sttt/metrics"
	"sync"
)

const Player1 = -1
const Player1Win = Player1 * 3
const Player2 = 1
const Player2Win = Player2 * 3

func NewBoard() *Board {
	return &Board{
		Tiles: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
		NextTurnID:   Player1, //TODO: randomize
		HasWinner:    false,
		WinnerTurnID: 0,
	}
}

type Board struct {
	sync.Mutex
	Tiles        [][]int `json:"tiles"`
	NextTurnID   int     `json:"nextTurnID"`
	HasWinner    bool    `json:"hasWinner"`
	WinnerTurnID int     `json:"winnerTurnID"`
	Started      bool    `json:"started"`
}

func (b *Board) Play(x, y, turnID int) error {
	b.Lock()
	defer b.Unlock()
	if b.HasWinner {
		return errors.New("You can't move after someone has won")
	}
	if turnID != b.NextTurnID {
		return errors.New("It's not your turn")
	}
	if b.Tiles[x][y] == 0 {
		b.Tiles[x][y] = turnID
		b.NextTurnID = b.NextTurnID * -1
		metrics.Plays.Inc()
		return nil
	}
	return errors.New(fmt.Sprintf("tile %v,%v is not 0", x, y))
}

func (b *Board) CheckForWinner() {
	b.Lock()
	defer b.Unlock()
	if b.HasWinner {
		return
	}
	defer func() {
		if b.HasWinner {
			metrics.Wins.Inc()
		}
	}()
	for row := range b.getRows() {
		if row == Player1Win {
			b.HasWinner = true
			b.WinnerTurnID = Player1
			break
		} else if row == Player2Win {
			b.HasWinner = true
			b.WinnerTurnID = Player2
			break
		}
	}
	return
}

func (b *Board) getRows() chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		//Horizontal
		result <- b.Tiles[0][0] + b.Tiles[0][1] + b.Tiles[0][2]
		result <- b.Tiles[1][0] + b.Tiles[1][1] + b.Tiles[1][2]
		result <- b.Tiles[2][0] + b.Tiles[2][1] + b.Tiles[2][2]
		//Vertical
		result <- b.Tiles[0][0] + b.Tiles[1][0] + b.Tiles[2][0]
		result <- b.Tiles[0][1] + b.Tiles[1][1] + b.Tiles[2][1]
		result <- b.Tiles[0][2] + b.Tiles[1][2] + b.Tiles[2][2]
		//Diagonal
		result <- b.Tiles[0][0] + b.Tiles[1][1] + b.Tiles[2][2]
		result <- b.Tiles[0][2] + b.Tiles[1][1] + b.Tiles[2][0]

	}()
	return result
}
