package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pelletier/go-toml/v2"
	// toml "repo_pusher/pkg/toml_parser"
)

type Config struct {
	Remote string
	Repos  []struct {
		Name string
		Path string
	}
}

func main() {
	config := ReadConfig("config.toml")
	PushChanges(&config)
}

func ReadConfig(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := toml.NewDecoder(file)
	config := Config{}

	err = decoder.Decode(&config)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	return config
}

func PushChanges(config *Config) {
	for i, repo := range config.Repos {
		cmd := exec.Command("git", "status", "--porcelain")
		cmd.Dir = repo.Path

		out, err := cmd.Output()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		} else {
			fmt.Printf("%d - %s: %v\n", i+1, repo.Name, string(out))
		}
	}
}
