package http_server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	// function without routine won't be called
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	// function without routine won't be called
	return s.league
}

func TestGetPlayers(t *testing.T) {
	cases := []struct {
		name string
		path string
		want string
	}{
		{"get Pepper's score", "/players/Pepper", "20"},
		{"get Floyd's score", "/players/Floyd", "10"},
	}

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, c.path, nil)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, http.StatusOK)

			got := response.Body.String()
			want := c.want
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}

	t.Run("non existant players score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Anon", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d, want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", "Pepper"), nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != "Pepper" {
			t.Errorf("got %q want %q", store.winCalls[0], "Pepper")
		}
	})
}

func TestLeague(t *testing.T) {

	t.Run("returns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)

		if !reflect.DeepEqual(got, wantedLeague) {
			t.Errorf("got %v want %v", got, wantedLeague)
		}

		if response.Result().Header.Get("Content-Type") != "application/json" {
			t.Errorf("got %v, want %v", got, "application/json")
		}
	})
}

func assertStatus(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("status %d, want %d", got, want)
	}
}
