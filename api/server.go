package api

import (
	"errors"
	"github.com/just1689/sttt/domain"
	"time"
)

var Server = domain.NewServer()
var maxTries = 20

func HandleQuickGame(p domain.Player) (game *domain.Game, err error) {
	tries := 0
	for tries < maxTries {
		tries++
		g := Server.NextOpenGame()
		if g == nil {
			continue
		}
		err = HandleJoinGame(g.ID, p)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		game = g
		return
	}
	game, err = HandleCreateGame(p)
	return
}

func HandleCreatePlayer(name string) *domain.Player {
	return Server.GeneratePlayer(name)
}

func HandleListGames() []domain.GameInfo {
	return Server.GetGames()
}

func HandleCreateGame(p domain.Player) (game *domain.Game, err error) {
	if p.ID == "" || p.Name == "" || p.Secret == "" {
		err = errors.New("missing player info")
		return
	}
	game = Server.CreateGame(&p)
	return
}

func HandleJoinGame(gameID string, player domain.Player) (err error) {
	game, found := Server.GameByID(gameID)
	if !found {
		err = errors.New("could not find game")
		return
	}
	err = game.NewPlayer(&player)
	return
}

func HandleGameInfo(gameID string) (game *domain.Game, err error) {
	var found bool
	game, found = Server.GameByID(gameID)
	if !found {
		err = errors.New("could not find game")
		return
	}
	return
}

func HandleTurn(gameID string, secret string, turnID int, x, y int) (err error) {
	game, found := Server.GameByID(gameID)
	if !found {
		err = errors.New("could not find gameID")
		return
	}
	player := game.PlayersInGame.GetPlayerByTurnID(turnID)
	if player == nil {
		err = errors.New("no player found at turnID")
		return
	}
	if player.Secret == secret {
		err = errors.New("secret does not match")
		return
	}
	err = game.Board.Play(x, y, turnID)
	return
}
