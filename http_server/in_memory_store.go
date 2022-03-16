package http_server

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	Store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.Store[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.Store[name]++
}
func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player
	for name, wins := range i.Store {
		league = append(league, Player{name, wins})
	}
	return league
}
