package http_server

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFileStorage(t *testing.T) {
	data := `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`
	t.Run("get league from file storage", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, data)
		defer cleanDatabase()
		store := FileStoragePlayerStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)

		got2 := store.GetLeague()
		assertLeague(t, got2, want)
	})
	t.Run("get score from file storage", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, data)
		defer cleanDatabase()
		store := FileStoragePlayerStore{database}
		// store := createStore(t, data)

		got := store.GetScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, data)
		defer cleanDatabase()
		store := FileStoragePlayerStore{database}
		store.RecordWin("Chris")
		got := store.GetScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})
}

func assertLeague(t testing.TB, got []Player, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertScoreEquals(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {

	t.Helper()
	tmpFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpFile.Write([]byte(initialData))
	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
	return tmpFile, removeFile
}

func createStore(t testing.TB, data string) (store FileStoragePlayerStore) {
	t.Helper()
	database, cleanDatabase := createTempFile(t, data)
	defer cleanDatabase()
	return FileStoragePlayerStore{database}
}
