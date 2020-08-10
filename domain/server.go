package domain

import (
	"github.com/google/uuid"
	"sync"
)

func NewServer() *Server {
	return &Server{
		Games: make(map[string]*Game),
	}
}

type Server struct {
	sync.Mutex
	Games map[string]*Game
}

func (s *Server) GetGames() []GameInfo {
	s.Lock()
	defer s.Unlock()
	r := make([]GameInfo, 0)
	for _, g := range s.Games {
		r = append(r, g.GetGameInfo())
	}
	return r
}

func (s *Server) CreateGame(p *Player) *Game {
	s.Lock()
	defer s.Unlock()
	g := NewGame(p)
	s.Games[g.ID] = g
	return g
}

func (s *Server) GeneratePlayer(name string) *Player {
	s.Lock()
	defer s.Unlock()
	p := &Player{
		ID:     uuid.New().String(),
		Secret: uuid.New().String(),
		Name:   name,
	}
	return p
}
