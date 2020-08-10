package domain

import (
	"errors"
	"sync"
)

func NewPlayersInGame(p *Player) *PlayersInGame {
	return &PlayersInGame{
		Players: map[int]*Player{
			Player1: p,
		},
	}
}

type PlayersInGame struct {
	sync.Mutex
	Players map[int]*Player `json:"list"`
}

func (p *PlayersInGame) AddPlayer(player *Player) (err error) {
	p.Lock()
	defer p.Unlock()
	if p.IsGameFull() {
		err = errors.New("game is full")
		return
	}
	_, found := p.Players[Player1]
	if found {
		p.Players[Player2] = player
		return
	}
	p.Players[Player1] = player
	return
}

func (p *PlayersInGame) RemovePlayer(player *Player) {
	p.Lock()
	defer p.Unlock()
	removeID := 0
	for id, pl := range p.Players {
		if pl.ID == player.ID {
			removeID = id
			break
		}
	}
	if removeID != 0 {
		delete(p.Players, removeID)
	}
	return
}

func (p *PlayersInGame) GetPlayerByID(playerID string) *Player {
	p.Lock()
	defer p.Unlock()
	for _, pl := range p.Players {
		if pl.ID == playerID {
			return pl
		}
	}
	return nil
}

func (p *PlayersInGame) GetPlayerByTurnID(id int) *Player {
	p.Lock()
	defer p.Unlock()
	return p.Players[id]
}

func (p *PlayersInGame) IsGameFull() bool {
	return len(p.Players) == 2
}
