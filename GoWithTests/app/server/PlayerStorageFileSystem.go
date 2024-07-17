package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type PlayerStorageFileSystem struct {
	database *json.Encoder
	league   League
}

func initDatabaseFilePlayer(storage *os.File) error {
	storage.Seek(0, 0)

	info, err := storage.Stat()

	if err != nil {
		return fmt.Errorf("could not stat file %s, %v", storage.Name(), err)
	}

	if info.Size() == 0 {
		storage.Write([]byte("[]"))
		storage.Seek(0, 0)
	}

	return nil
}
func NewPlayerStorageFileSystem(storage *os.File) (*PlayerStorageFileSystem, error) {
	storage.Seek(0, io.SeekStart)

	err := initDatabaseFilePlayer(storage)
	if err != nil {
		return nil, fmt.Errorf("error during start file %s, %v", storage.Name(), err)
	}
	league, err := NewLeague(storage)

	if err != nil {
		return nil, fmt.Errorf("error loading Storage: %v", err)
	}

	return &PlayerStorageFileSystem{
		database: json.NewEncoder(&tape{storage}),
		league:   league,
	}, nil
}

func (f *PlayerStorageFileSystem) GetLeague() League {
	/*	_, _ = f.database.Seek(0, io.SeekStart)
		league, _ := NewLeague(f.database)
		return league*/
	return f.league
}

func (f *PlayerStorageFileSystem) Get(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *PlayerStorageFileSystem) Save(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	/*	_, _ = f.database.Seek(0, io.SeekStart)*/
	/*	_ = json.NewEncoder(f.database.).Encode(f.league)*/
	f.database.Encode(f.league)
}
