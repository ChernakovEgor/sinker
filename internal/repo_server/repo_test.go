package repo_server

import (
	"testing"

	cr "github.com/ChernakovEgor/sinker/internal/config_reader"
)

type StubShell struct {
}

func (s StubShell) Execute(cmd string) ([]byte, error) {
	return nil, nil
}

// TODO
func TestGetUpdates(t *testing.T) {
	serverConfig := cr.ServerConfig{Name: "mock", Remote: "", Host: "", Port: 1}
	var shell Shell = StubShell{}
	var server RepoServer = NewSSHServer(shell, serverConfig)

	want := 4
	got, err := server.FetchUpdates()
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if got != want {
		t.Errorf("Number of updates is wrong, got %d want %d", got, want)
	}
}
