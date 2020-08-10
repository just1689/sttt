package web

import "net/http"

func Setup(listen string) {
	http.HandleFunc("/createPlayer", handleCreatePlayer)
	http.HandleFunc("/listGames", handleListGames)
	http.HandleFunc("/createGame", handleCreateGame)
	http.HandleFunc("/joinGame", handleJoinGame)
	http.HandleFunc("/turn", handleTurn)
	http.HandleFunc("/gameInfo", HandleGetGameInfo)
	http.ListenAndServe(listen, nil)
}
