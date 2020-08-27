package web

import (
	"github.com/just1689/sttt/backend/metrics"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Setup(listen string) {
	http.HandleFunc("/sttt/api/createPlayer", wrapper(handleCreatePlayer))
	http.HandleFunc("/sttt/api/listGames", wrapper(handleListGames))
	http.HandleFunc("/sttt/api/createGame", wrapper(handleCreateGame))
	http.HandleFunc("/sttt/api/joinGame", wrapper(handleJoinGame))
	http.HandleFunc("/sttt/api/turn", wrapper(handleTurn))
	http.HandleFunc("/sttt/api/gameInfo", wrapper(HandleGetGameInfo))
	http.HandleFunc("/sttt/api/quickGame", wrapper(handleQuickGame))
	http.HandleFunc("/", wrapper(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(request.URL.RawPath))
	}))

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listen, nil)
}

func wrapper(f func(writer http.ResponseWriter, request *http.Request)) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		logrus.Infoln("HTTP request:", request.URL.Path)
		metrics.HTTPCalls.Inc()
		f(writer, request)
	}
}
