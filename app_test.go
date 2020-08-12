package main

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/sttt/api"
	"sync"
	"testing"
	"time"
)

func Test_endToEndMultiThreaded(t *testing.T) {
	api.RunQuickGame()

	time.Sleep(1 * time.Second)

	games := 2000
	wg := sync.WaitGroup{}
	for i := 0; i < games; i++ {
		wg.Add(1)
		go func() {
			p1N := fmt.Sprint("x", i)
			p2N := fmt.Sprint("y", i)
			tryTwoPlayersAsync2(p1N, p2N, t)
			wg.Done()

		}()
	}
	wg.Wait()
	count := len(api.Server.Games)
	if count > games {
		t.Error("bad # of games ", count, " instead of ", games)
	}

}

func tryTwoPlayers(p1Name, p2Name string, t *testing.T) {
	p1 := api.HandleCreatePlayer(p1Name)
	p2 := api.HandleCreatePlayer(p2Name)

	game := api.HandleQuickGame(*p1)
	if game == nil {
		t.Error("could not quick join ", p1Name)
		return
	}

	p2Game := api.HandleQuickGame(*p2)
	if p2Game == nil {
		t.Error("could not quick join ", p2Name)
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

func tryTwoPlayersAsync2(p1Name, p2Name string, t *testing.T) {
	p1 := api.HandleCreatePlayer(p1Name)
	p2 := api.HandleCreatePlayer(p2Name)

	game := api.HandleQuickGame(*p1)
	if game == nil {
		t.Error("could not quick join ", p1Name)
		return
	}

	game2 := api.HandleQuickGame(*p2)
	if game2 == nil {
		t.Error("could not quick join ", p2Name)
		return
	}

	_, err := api.HandleGameInfo(game.ID)
	if err != nil {
		t.Error(err)
		return
	}

}
