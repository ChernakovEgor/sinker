package repo

type RepoServer struct {
}

func (r RepoServer) FetchUpdates() (num int, err error) {
	num = 4
	return
}
