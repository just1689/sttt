package web

import "github.com/just1689/sttt/backend/domain"

type createPlayerRequest struct {
	Name string `json:"name"`
}

type joinGameRequest struct {
	GameID string        `json:"gameID"`
	Player domain.Player `json:"player"`
}

type gameInfoRequest struct {
	GameID string `json:"gameID"`
}

type turnRequest struct {
	GameID string `json:"gameID"`
	Secret string `json:"secret"`
	TurnID int    `json:"turnID"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}
