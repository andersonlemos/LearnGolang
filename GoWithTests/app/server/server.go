package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const contentTypeJSON = "application/json"

type Player struct {
	Name string
	Wins int
}

type PlayerStorage interface {
	GetPlayerScore(name string) int
	AddWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	storage PlayerStorage
	http.Handler
}

func NewPlayerServer(storage PlayerStorage) *PlayerServer {
	s := new(PlayerServer)

	s.storage = storage

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.handleLeague))
	router.Handle("/players/", http.HandlerFunc(s.handlePlayer))

	s.Handler = router

	return s
}

/*
	func (s *PlayerServer) getTableLeague() []Player {
		return []Player{
			{"Chris", 20},
		}
	}
*/
func (s *PlayerServer) handleLeague(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", contentTypeJSON)
	_ = json.NewEncoder(w).Encode(s.storage.GetLeague())
}
func (s *PlayerServer) handlePlayer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		s.showScore(w, player)
	case http.MethodPost:
		s.AddWin(w, player)
	}
}

func (s *PlayerServer) GetPlayerScore(name string) int {
	switch name {
	case "Mary":
		return 20
	case "Peter":
		return 10
	default:
		return 0
	}
}

func (s *PlayerServer) AddWin(w http.ResponseWriter, player string) {
	s.storage.AddWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (s *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := s.storage.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprint(w, s.storage.GetPlayerScore(player))
}
