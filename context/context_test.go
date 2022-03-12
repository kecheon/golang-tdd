package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	data := "Hello world"
	stubStore := &StubStore{response: data, cancelled: false}
	svr := Server(stubStore)
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(500*time.Millisecond, cancel)
	request = request.WithContext(cancellingCtx)
	svr.ServeHTTP(response, request)
	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
	if stubStore.cancelled {
		t.Errorf("should not be cancelled")
	}
}

type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}
func (s *StubStore) Cancel() {
	s.cancelled = true
}
