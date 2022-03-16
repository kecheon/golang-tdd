package http_server

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileStoragePlayerStore struct {
	database io.ReadWriteSeeker
}

func (s *FileStoragePlayerStore) GetLeague() League {
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

func (s *FileStoragePlayerStore) GetScore(name string) int {
	player := s.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (s *FileStoragePlayerStore) RecordWin(name string) {
	league := s.GetLeague()
	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
			/* TOD
			reason not using player.Wins++ because
			range over the slice returns copy of the element
			updating Wins value of a copy element has no effect
			have to update the original league
			*/
			break
		}
	}
	s.database.Seek(0, 0)
	json.NewEncoder(s.database).Encode(league)
}
