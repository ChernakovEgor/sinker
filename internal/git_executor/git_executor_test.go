package git_executor

import (
	"strings"
	"testing"

	cr "github.com/ChernakovEgor/sinker/internal/config_reader"
)

type StubShell struct {
}

func (s StubShell) Execute(cmd string) (string, error) {
	commands := strings.Split(cmd, " ")
	switch commands[1] {
	case "status":
		return Status(commands[2])
	}
	return "", nil
}

// TODO: write status tests
func Status(repo string) (string, error) {
	switch repo {
	case "repo_1":
		return "m", nil
	}
	return "", nil
}

func TestStatus(t *testing.T) {
	repos := []cr.RepoConfig{
		{Name: "repo_1", Path: "~/dev/repo_1"},
		{Name: "repo_2", Path: "~/dev/repo_2"},
		{Name: "repo_3", Path: "~/dev/repo_3"},
		{Name: "repo_4", Path: "~/dev/repo_4"},
	}

	want := `repo_1 1M 1D 1?
  repo_2 3M 2?
  repo_3 ok
  repo_4 not a repo`

	gitShell := StubShell{}
	git := NewLocalGit(repos, gitShell)

	got, _ := git.Status()
	if want != got {
		t.Errorf("invalid output: got %s want %s", got, want)
	}
}
