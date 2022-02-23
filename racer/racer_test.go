package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacerWithMock(t *testing.T) {
	slow := makeServer(20 * time.Millisecond)
	fast := makeServer(0 * time.Millisecond)
	defer slow.Close()
	defer fast.Close()

	got := Racer(slow.URL, fast.URL)
	want := fast.URL

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}

}

func makeServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacerPing(t *testing.T) {
	t.Run("get fast url", func(t *testing.T) {
		slow := makeServer(20 * time.Millisecond)
		fast := makeServer(0 * time.Millisecond)
		defer slow.Close()
		defer fast.Close()

		got, _ := RacerPing(slow.URL, fast.URL, 100*time.Millisecond)
		want := fast.URL

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("error if no response in 10 seconds", func(t *testing.T) {
		slow := makeServer(25 * time.Millisecond)
		defer slow.Close()

		_, err := RacerPing(slow.URL, slow.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("exected error but didn't get one")
		}
	})
}
