package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPCalls = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sttt_http_calls",
		Help: "The total number of http calls",
	})

	GamesCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sttt_games_created",
		Help: "The total number of games created",
	})

	PlayersJoined = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sttt_players_joined",
		Help: "The total number of players joined",
	})

	Plays = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sttt_players_plays",
		Help: "The total number of plays made",
	})

	Wins = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sttt_wins",
		Help: "The total number of wins",
	})
)
