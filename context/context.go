package context

import (
	"fmt"
	"net/http"
)

func Server(s Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)
		go func() {
			data <- s.Fetch()
		}()
		select {
		case d := <-data:
			fmt.Fprintf(w, d)
		case <-ctx.Done():
			s.Cancel()
		}
	}
}

type Store interface {
	Fetch() string
	Cancel()
}
