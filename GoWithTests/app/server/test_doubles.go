package server

type SketchPlayerStorage struct {
	scores    map[string]int
	winScores []string
	league    League
}

func (s *SketchPlayerStorage) GetPlayerScore(name string) int {
	return s.scores[name]
}
func (s *SketchPlayerStorage) AddWin(name string) {
	s.winScores = append(s.winScores, name)
}
func (s *SketchPlayerStorage) GetLeague() League {
	return s.league
}
