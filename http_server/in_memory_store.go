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
