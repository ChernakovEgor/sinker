package connector

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Connection interface {
	Ping() bool
}

type SSHConnection struct {
	host string
	user string
	port string
}

func CreateSSHConnection(host, user, port string) *SSHConnection {
	return &SSHConnection{host, user, port}
}

func (s SSHConnection) Ping() (ok bool) {
	cmd := exec.Command("ssh", s.host, "-p", s.port, "bla")
	var errBytes bytes.Buffer
	cmd.Stderr = &errBytes

	cmd.Run()
	if errBytes.Len() != 0 {
		fmt.Printf("ping error: %v\n", string(errBytes.Bytes()))
		return false
	}

	return true
}
