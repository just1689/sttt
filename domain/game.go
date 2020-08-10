package domain

import (
	"github.com/google/uuid"
)

func NewGame(p *Player) *Game {
	return &Game{
		ID:            uuid.New().String(),
		Board:         NewBoard(),
		PlayersInGame: NewPlayersInGame(p),
	}
}

type Game struct {
	ID            string         `json:"game"`
	Board         *Board         `json:"board"`
	PlayersInGame *PlayersInGame `json:"players"`
}

func (g *Game) NewPlayer(p *Player) (err error) {
	if err = g.PlayersInGame.AddPlayer(p); err != nil {
		return
	}
	if g.PlayersInGame.IsGameFull() {
		g.Board.Started = true
	}
	return
}

func (g *Game) PlayerLeft(p *Player) {
	g.PlayersInGame.RemovePlayer(p)
}

func (g *Game) GetGameInfo() GameInfo {
	return GameInfo{
		ID:    g.ID,
		Board: g.Board,
		Full:  g.PlayersInGame.IsGameFull(),
	}
}

//GameInfo is used to list games
type GameInfo struct {
	ID    string `json:"id"`
	Board *Board `json:"board"`
	Full  bool   `json:"full"`
}
