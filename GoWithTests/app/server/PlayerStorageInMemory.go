package server

type PlayerStorageInMemory struct {
	storage map[string]int
}

func (s *PlayerStorageInMemory) GetLeague() League {
	var league []Player
	for name, wins := range s.storage {
		league = append(league, Player{name, wins})
	}
	return league
}

func (s *PlayerStorageInMemory) AddWin(player string) {
	s.storage[player]++
}

func (s *PlayerStorageInMemory) GetPlayerScore(player string) int {
	return s.storage[player]
}
