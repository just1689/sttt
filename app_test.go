package main

import (
	"encoding/json"
	"github.com/just1689/sttt/api"
	"testing"
)

func Test_endToEnd2Players(t *testing.T) {
	tryTwoPlayers("p1", "p2", t)

}

func Test_endToEnd4Players(t *testing.T) {
	tryTwoPlayers("p1", "p2", t)
	tryTwoPlayers("p3", "p4", t)

}

func tryTwoPlayers(p1Name, p2Name string, t *testing.T) {
	p1 := api.HandleCreatePlayer(p1Name)
	p2 := api.HandleCreatePlayer(p2Name)

	game, err := api.HandleQuickGame(*p1)
	if err != nil {
		t.Error("could not quick join ", p1Name, err)
		return
	}

	p2Game, err := api.HandleQuickGame(*p2)
	if err != nil {
		t.Error("could not quick join ", p2Name, err)
		return
	}

	if game.ID != p2Game.ID {
		t.Error("Two players join but got two different games!")
		return
	}

	info, err := api.HandleGameInfo(game.ID)
	if err != nil {
		t.Error(err)
		return
	}

	b, _ := json.Marshal(info)
	t.Log(string(b))
}
