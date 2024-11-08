package git_executor

type GitExecutor interface {
	FetchUpdates()
	Push()
	Status() (string, error)
}
