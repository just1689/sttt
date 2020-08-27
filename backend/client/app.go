package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/sttt/backend/domain"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var baseURL = "https://lab.captainjustin.space/sttt/api/"

func main() {

	humansCount := 1000
	wg := sync.WaitGroup{}

	for i := 0; i < humansCount; i++ {
		wg.Add(1)
		RunHuman(10, &wg)
		time.Sleep(100 * time.Millisecond)
	}
	go func() {
		for {
			fmt.Println("Busy...")
			time.Sleep(3 * time.Second)
		}
	}()
	wg.Wait()

}

func RunHuman(totalGames int, wg *sync.WaitGroup) {
	go func() {
		myDetails := domain.Auth{
			ID:     uuid.New().String(),
			Secret: uuid.New().String(),
			Name:   "",
		}
		myDetails.Name = myDetails.ID
		var game *domain.Game
		for totalGames > 0 {

			//Make sure there is a game
			if game == nil {
				b, err := post("quickGame", myDetails)
				if err != nil {
					logrus.Errorln("could not quick join")
					time.Sleep(1 * time.Second)
					continue
				}
				totalGames--
				game = gameFromBytes(b)
				if game == nil {
					logrus.Errorln("could unmarshal game")
					logrus.Errorln(string(b))
					time.Sleep(1 * time.Second)
					continue
				}
			}

			//Play when it is my turn

			//Pretend the game is over
			time.Sleep(2 * time.Second)
			game = nil
			continue

		}
		wg.Done()
	}()
}

var JsonContent = "Application/json"

func post(api string, i interface{}) (b []byte, err error) {
	var bi []byte
	bi, err = json.Marshal(i)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	resp, err := http.Post(fmt.Sprint(baseURL, api), JsonContent, bytes.NewReader(bi))
	if err != nil {
		logrus.Error(err)
		return
	}
	b, err = ioutil.ReadAll(resp.Body)
	return

}

func gameFromBytes(b []byte) *domain.Game {
	result := &domain.Game{}
	err := json.Unmarshal(b, result)
	if err != nil {
		logrus.Errorln("could not unmarshal game")
	}
	return result
}
