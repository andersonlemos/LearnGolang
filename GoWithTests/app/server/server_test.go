package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	/*storage := SketchPlayerStorage{
		scores:    map[string]int{"Mary": 20, "Peter": 10},
		winScores: nil,
	}*/

	database, cleanDatabase := createTemporaryFile(t, "")
	defer cleanDatabase()

	ps := NewPlayerServer(&database)

	t.Run("return Mary's results", func(t *testing.T) {
		request := newRequestToGetPoints("Mary")
		response := httptest.NewRecorder()

		ps.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkResponse(t, response.Body.String(), "20")

	})

	t.Run("returns Peter's results", func(t *testing.T) {
		request := newRequestToGetPoints("Peter")
		response := httptest.NewRecorder()

		ps.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkResponse(t, response.Body.String(), "10")
	})

	t.Run("return 404 if player not found", func(t *testing.T) {
		request := newRequestToGetPoints("Josh")
		response := httptest.NewRecorder()

		ps.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusNotFound)

	})
}

func TestWinStorage(t *testing.T) {

	t.Run("returns Status 'ACCEPTED' to POST calls", func(t *testing.T) {
		storage := SketchPlayerStorage{
			scores:    map[string]int{},
			winScores: nil,
		}

		server := NewPlayerServer(&storage)

		request := newRequestToAddWin("Mary")

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		checkStatusCode(t, response.Code, http.StatusAccepted)
	})
	t.Run("should register wins on HTTP POST calls", func(t *testing.T) {
		storage := SketchPlayerStorage{
			scores:    map[string]int{},
			winScores: nil,
		}

		server := NewPlayerServer(&storage)
		player := "Mary"
		request := newRequestToAddWin(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusAccepted)

		if len(storage.winScores) != 1 {
			t.Errorf("Expected 1 player wins, got %d", len(storage.winScores))
		}
		if storage.winScores[0] != player {
			t.Errorf("Expected player '%s' to be wins, got '%s'", player, storage.winScores[0])
		}
	})
}

func TestRegisterVictoryAndSearch(t *testing.T) {
	database, cleanDatabase := createTemporaryFile(t, "")
	defer cleanDatabase()
	storage := &PlayerStorageFileSystem{database: database}
	server := NewPlayerServer(storage)
	player := "Mary"

	server.ServeHTTP(httptest.NewRecorder(), newRequestToAddWin(player))
	server.ServeHTTP(httptest.NewRecorder(), newRequestToAddWin(player))
	server.ServeHTTP(httptest.NewRecorder(), newRequestToAddWin(player))

	response := httptest.NewRecorder()

	server.ServeHTTP(response, newRequestToGetPoints(player))

	checkStatusCode(t, response.Code, http.StatusOK)

	checkResponse(t, response.Body.String(), "3")
}

func TestLeague(t *testing.T) {

	t.Run("returns 200 in /league", func(t *testing.T) {
		storage := SketchPlayerStorage{}
		server := NewPlayerServer(&storage)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var players []Player

		err := json.NewDecoder(response.Body).Decode(&players)

		if err != nil {
			t.Fatalf("Unable to decode server '%s' from player: '%v'", response.Body, err)
		}
		checkStatusCode(t, response.Code, http.StatusOK)
	})
	t.Run("should return JSON from Table League", func(t *testing.T) {
		expectedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		storage := SketchPlayerStorage{nil, nil, expectedLeague}
		server := NewPlayerServer(&storage)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueResponse(t, response.Body)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkLeagueResponse(t, got, expectedLeague)
		checkContentType(t, response, contentTypeJSON)

	})
}

func TestPlayerStorageFileSystem(t *testing.T) {
	t.Run("/reader's league", func(t *testing.T) {
		/*		database := strings.NewReader(`[
				{"Name": "Cleo","Wins": 10},
				{"Name": "Chris","Wins": 33}]`)
		*/
		database, cleanupDatabase := createTemporaryFile(t, `[
		{"Name": "Cleo","Wins": 10},
		{"Name": "Chris","Wins": 33}]`)

		defer cleanupDatabase()

		storage := PlayerStorageFileSystem{database}
		got := storage.GetLeague()

		expected := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		checkLeagueResponse(t, got, expected)

	})

	t.Run("get Player's score", func(t *testing.T) {
		database, cleanupDatabase := createTemporaryFile(t, `[
		{"Name": "Cleo","Wins": 10},
		{"Name": "Chris","Wins": 33}]`)

		defer cleanupDatabase()

		storage := PlayerStorageFileSystem{database}

		got := storage.Get("Chris")
		expected := 33
		checkEqualScores(t, got, expected)
	})

	t.Run("stores wins of an existing player", func(t *testing.T) {
		database, cleanupDatabase := createTemporaryFile(t, `[
		{"Name": "Cleo","Wins": 10},
		{"Name": "Chris","Wins": 33}]`)
		defer cleanupDatabase()

		storage := PlayerStorageFileSystem{database}

		storage.Save("Chris")
		got := storage.Get("Chris")
		expected := 34
		checkEqualScores(t, got, expected)
	})
	t.Run("save new players win", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := createTemporaryFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer limpaBancoDeDados()

		storage := PlayerStorageFileSystem{bancoDeDados}

		storage.Save("Pepper")

		got := storage.Get("Pepper")
		expected := 1
		checkEqualScores(t, got, expected)
	})
}

func checkEqualScores(t *testing.T, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("got '%d' want '%d'", got, expected)
	}
}

func createTemporaryFile(t *testing.T, data string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatal("Unable to create temporary file")
	}

	_, _ = tmpFile.Write([]byte(data))

	removeFile := func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}
	return tmpFile, removeFile
}

func checkLeagueResponse(t *testing.T, got, expected []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func checkContentType(t *testing.T, response *httptest.ResponseRecorder, expected string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != expected {
		t.Errorf("response does not contain content-type from %s, got %v", expected, response.Result().Header)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newRequestToGetPoints(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func newRequestToAddWin(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func checkResponse(t *testing.T, receive, expected string) {
	t.Helper()
	if receive != expected {
		t.Errorf("Response body doesn't match. Expected '%s', got '%s'", expected, receive)
	}
}

func checkStatusCode(t *testing.T, receive, expected int) {
	t.Helper()

	if receive != expected {
		t.Errorf("HTTP Status Codes doesn't match. Expected '%d', got '%d'", expected, receive)
	}
}

func getLeagueResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("error decoding server response '%s': %v", body, err)
	}

	return
}
