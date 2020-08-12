package main

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/sttt/api"
	"sync"
	"testing"
)

func Test_endToEnd2Players(t *testing.T) {
	tryTwoPlayers("p1", "p2", t)

}

func Test_endToEnd4Players(t *testing.T) {
	tryTwoPlayers("p1", "p2", t)
	tryTwoPlayers("p3", "p4", t)

}

func Test_endToEndMultiThreaded(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			p1N := fmt.Sprint("x", i)
			p2N := fmt.Sprint("y", i)
			tryTwoPlayersAsync(p1N, p2N, t)
			wg.Done()

		}()
	}
	wg.Wait()

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

func tryTwoPlayersAsync(p1Name, p2Name string, t *testing.T) {
	p1 := api.HandleCreatePlayer(p1Name)
	p2 := api.HandleCreatePlayer(p2Name)

	game, err := api.HandleQuickGame(*p1)
	if err != nil {
		t.Error("could not quick join ", p1Name, err)
		return
	}

	_, err = api.HandleQuickGame(*p2)
	if err != nil {
		t.Error("could not quick join ", p2Name, err)
		return
	}

	_, err = api.HandleGameInfo(game.ID)
	if err != nil {
		t.Error(err)
		return
	}

}
