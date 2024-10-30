package main

import (
	// "fmt"
	"fmt"
	reader "github.com/ChernakovEgor/sinker/internal/config_reader"
	"os"
	"os/exec"
)

func main() {
	config, err := reader.ReadConfig("config.toml")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	cmd := exec.Command

	fmt.Println(*config)
}
