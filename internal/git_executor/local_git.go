package git_executor

import configreader "github.com/ChernakovEgor/sinker/internal/config_reader"

type Shell interface {
	Execute(cmd string) (string, error)
}

type LocalGit struct {
	repos []configreader.RepoConfig
	shell Shell
}

func (l LocalGit) Status() (string, error) {
	result := `repo_1 1M 1D 1?
  repo_2 3M 2?
  repo_3 ok
  repo_4 not a repo`
	return result, nil
}

func (l LocalGit) FetchUpdates() {
}

func (l LocalGit) Push() {
}

func NewLocalGit(repos []configreader.RepoConfig, shell Shell) GitExecutor {
	return LocalGit{}
}
