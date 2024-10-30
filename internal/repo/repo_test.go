package repo

import (
	"testing"
)

func TestGetUpdates(t *testing.T) {
	server := RepoServer{}
	num, err := server.FetchUpdates()

	want := 4
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if num != want {
		t.Errorf("Number of updates is wrong, got %d want %d", num, want)
	}
}
