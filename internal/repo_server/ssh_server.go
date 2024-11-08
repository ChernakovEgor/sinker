package repo_server

import (
	cr "github.com/ChernakovEgor/sinker/internal/config_reader"
	"os/exec"
)

type Shell interface {
	Execute(cmd string) ([]byte, error)
}

// LocalShell - default shell executor
type LocalShell struct {
}

func (_ LocalShell) Execute(cmd string) ([]byte, error) {
	// localShell := exec.CommandContext(ctx, cmd)
	localShell := exec.Command(cmd)
	return localShell.CombinedOutput()
}

// SSH Server is used to connect to git server via SSH
type SSHServer struct {
	shell Shell
}

func NewSSHServer(shell Shell, config cr.ServerConfig) SSHServer {
	sshServer := SSHServer{shell: shell}
	return sshServer
}

func (s SSHServer) FetchUpdates() (int, error) {
	return 4, nil
}
