package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/just1689/sttt/api"
	"github.com/just1689/sttt/domain"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func handleCreatePlayer(writer http.ResponseWriter, request *http.Request) {
	r := &createPlayerRequest{}
	b, err := readBody(request)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - json", http.StatusBadRequest)
		return
	}
	p := api.HandleCreatePlayer(r.Name)
	a := p.GetAuth()
	marshalReply(writer, a)
}

func handleListGames(writer http.ResponseWriter, request *http.Request) {
	l := api.HandleListGames()
	marshalReply(writer, l)
}

func handleCreateGame(writer http.ResponseWriter, request *http.Request) {
	r := &domain.Auth{}
	b, err := readBody(request)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - json", http.StatusBadRequest)
		return
	}
	game, err := api.HandleCreateGame(r.GetPlayer())
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, fmt.Sprint("bad request - ", err.Error()), http.StatusBadRequest)
		return
	}
	marshalReply(writer, game)

}

func handleQuickGame(writer http.ResponseWriter, request *http.Request) {
	r := &domain.Auth{}
	b, err := readBody(request)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - json", http.StatusBadRequest)
		return
	}
	game := api.HandleQuickGame(r.GetPlayer())
	if game == nil {
		logrus.Errorln(errors.New("internal error quick joining game"))
		http.Error(writer, fmt.Sprint("bad request - ", err), http.StatusBadRequest)
		return
	}
	marshalReply(writer, game)

}

func handleJoinGame(writer http.ResponseWriter, request *http.Request) {
	r := &joinGameRequest{}
	b, err := readBody(request)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - json", http.StatusBadRequest)
		return
	}
	err = api.HandleJoinGame(r.GameID, r.Player)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := resultResponse{OK: true}
	marshalReply(writer, result)
}

func HandleGetGameInfo(writer http.ResponseWriter, request *http.Request) {
	r := &gameInfoRequest{}
	b, err := readBody(request)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - json", http.StatusBadRequest)
		return
	}
	game, err := api.HandleGameInfo(r.GameID)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	marshalReply(writer, game)
}

func handleTurn(writer http.ResponseWriter, request *http.Request) {
	r := &turnRequest{}
	b, err := readBody(request)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, r)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, "bad request - json", http.StatusBadRequest)
		return
	}
	err = api.HandleTurn(r.GameID, r.Secret, r.TurnID, r.X, r.Y)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := resultResponse{OK: true}
	marshalReply(writer, result)

}

func readBody(request *http.Request) (b []byte, err error) {
	defer request.Body.Close()
	b, err = ioutil.ReadAll(request.Body)
	return
}

func marshalReply(writer http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		logrus.Errorln(err)
		http.Error(writer, fmt.Sprint("internal error - json", err.Error()), http.StatusInternalServerError)
		return
	}
	reply(writer, b)
}

func reply(writer http.ResponseWriter, b []byte) {
	writer.Header().Add("content-type", "application/json")
	writer.Write(b)
}
