package load

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

type Connection interface {
	Ping() bool
}

type SSHConnection struct {
	host string
	port string
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

func ReadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var conn Connection = SSHConnection{"git@egorgeous.xyz", "903"}
	if !conn.Ping() {
		return fmt.Errorf("could not connect via ssh\n")
	}
	return nil
}
