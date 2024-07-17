package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRunner(t *testing.T) {
	t.Run("should return an error if the server not responding in 10 seconds", func(t *testing.T) {
		serverA := createServerWithDelay(25 * time.Second)

		defer serverA.Close()

		_, err := ConfigurableRunner(serverA.URL, serverA.URL, 20*time.Second)

		if err == nil {
			t.Error("expected error to be nil")
		}
	})
	t.Run("should compare the servers speed and returns the faster ", func(t *testing.T) {
		slowServer := createServerWithDelay(20 * time.Second)
		fastServer := createServerWithDelay(0 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		expected := fastServer.URL
		result, err := Runner(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("not waiting for an error but got %s", err)
		}
		if result != expected {
			t.Errorf("expected result to be %s but was %s", expected, result)
		}
	})
}
