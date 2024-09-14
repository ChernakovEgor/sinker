package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pelletier/go-toml/v2"
	// toml "repo_pusher/pkg/toml_parser"
)

const usage = `
usage: rp COMMAND

Commands:
  pull: pull updates from remotes
  push: push updates to remotes
`

type Config struct {
	Remote string
	Repos  []struct {
		Name string
		Path string
	}
}

func main() {

	if len(os.Args) < 2 || (os.Args[1] != "pull" && os.Args[1] != "push") {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "pull":
		config := ReadConfig("config.toml")
		PullUpdated(&config)
	case "push":
		config := ReadConfig("config.toml")
		PushChanges(&config)
	default:
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
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

func PullUpdated(config *Config) {
	for i, repo := range config.Repos {
		cmd := exec.Command("git", "fetch", "--all")
		cmd.Dir = repo.Path

		out, err := cmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%d - %s: %v\n", i+1, repo.Name, err)
		} else {
			fmt.Printf("%d - %s: %v\n", i+1, repo.Name, string(out))
		}
	}
}

func PushChanges(config *Config) {
	for i, repo := range config.Repos {
		cmd := exec.Command("git", "status", "--porcelain")
		cmd.Dir = repo.Path

		out, err := cmd.Output()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		} else {
			if len(out) > 0 {
				fmt.Printf("%d - %s: %v\n", i+1, repo.Name, "uncommited changes")
			} else {
				cmd = exec.Command("git", "push", "--all")
				cmd.Dir = repo.Path
				out, _ := cmd.Output()
				fmt.Printf("%d - %s: %v\n", i+1, repo.Name, string(out))
			}
		}
	}
}
