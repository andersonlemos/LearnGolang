package _context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	var data = "Hello, World!"

	t.Run("Returns store data", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("expected response body to be %q, got %q", data, response.Body.String())
		}

	})
	t.Run("Notifies the store to cancel the job if there is a cancellation request", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())

		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("expected response body to not be written")
		}

	})
}
