package http_server

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileStoragePlayerStore struct {
	database io.ReadSeeker
}

func (s *FileStoragePlayerStore) GetLeague() []Player {
	s.database.Seek(0, 0)
	league, _ := NewLeague(s.database)
	return league
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league %v", err)
	}
	return league, err
}

func (s *FileStoragePlayerStore) GetScore(name string) (int, error) {
	var wins int
	for _, player := range s.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins, nil
}
