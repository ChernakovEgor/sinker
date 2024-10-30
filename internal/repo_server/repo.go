package repo_server

import (
// cr "github.com/ChernakovEgor/sinker/internal/config_reader"
)

type RepoServer interface {
	FetchUpdates() (int, error)
}
