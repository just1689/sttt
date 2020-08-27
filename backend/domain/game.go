package domain

import (
	"github.com/google/uuid"
	"github.com/just1689/sttt/backend/metrics"
	"github.com/sirupsen/logrus"
)

func NewGame(p *Player) *Game {
	logrus.Infoln("Creating new game...")
	metrics.GamesCreated.Inc()
	metrics.PlayersJoined.Inc()
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
	metrics.PlayersJoined.Inc()
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
