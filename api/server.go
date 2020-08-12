package api

import (
	"errors"
	"github.com/just1689/sttt/domain"
	"github.com/sirupsen/logrus"
)

var Server = domain.NewServer()
var quickIn = make(chan Quick, 64)

func HandleQuickGame(p domain.Player) *domain.Game {
	reply := make(chan *domain.Game)
	quickIn <- Quick{
		p:     p,
		reply: reply,
	}
	return <-reply
}

type Quick struct {
	p     domain.Player
	reply chan *domain.Game
}

func RunQuickGame() {
	go func() {
		for next := range quickIn {
			done := false
			for !done {
				g := Server.NextOpenGame()
				if g != nil {
					err := HandleJoinGame(g.ID, next.p)
					if err != nil {
						done = false
						continue
					}
					next.reply <- g
					close(next.reply)
					done = true
					break
				}
				game, err := HandleCreateGame(next.p)
				if err != nil {
					logrus.Errorln("could not create game")
					continue
				}
				next.reply <- game
				close(next.reply)
				done = true
				break
			}
		}
	}()
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
