package web

import (
	"github.com/just1689/sttt/metrics"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Setup(listen string) {
	http.HandleFunc("/createPlayer", wrapper(handleCreatePlayer))
	http.HandleFunc("/listGames", wrapper(handleListGames))
	http.HandleFunc("/createGame", wrapper(handleCreateGame))
	http.HandleFunc("/joinGame", wrapper(handleJoinGame))
	http.HandleFunc("/turn", wrapper(handleTurn))
	http.HandleFunc("/gameInfo", wrapper(HandleGetGameInfo))
	http.HandleFunc("/quickGame", wrapper(handleQuickGame))

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listen, nil)
}

func wrapper(f func(writer http.ResponseWriter, request *http.Request)) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		metrics.HTTPCalls.Inc()
		f(writer, request)
	}
}
