package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := StubPlayserStore{
		map[string]int{
			"pepper": 20,
			"floyd":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("get Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/floyd", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "10"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

type StubPlayserStore struct {
	scores map[string]int
}

func (s *StubPlayserStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
