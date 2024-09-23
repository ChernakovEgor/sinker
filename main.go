package main

import (
	"bytes"
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
		Name        string
		Path        string
		Auto_Commit bool
		Auto_merge  bool
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
		var errb bytes.Buffer
		fetch := exec.Command("git", "fetch", "--all")
		fetch.Dir = repo.Path
		fetch.Stderr = &errb

		out, err := fetch.Output()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error %d - %s: %v\n", i+1, repo.Name, errb.String())
			continue
		}

		if repo.Auto_merge {
			merge := exec.Command("git", "pull", "dev")
			merge.Dir = repo.Path
			merge.Stderr = &errb

			out, err = merge.Output()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error %d - %s: %s\n", i+1, repo.Name, errb.String())
				continue
			}
		}
		fmt.Printf("%d - %s: %v\n", i+1, repo.Name, string(out))
	}
}

func PushChanges(config *Config) {
	for i, repo := range config.Repos {
		var errb bytes.Buffer
		status := exec.Command("git", "status", "--porcelain")
		status.Dir = repo.Path
		status.Stderr = &errb

		out, err := status.Output()

		if err != nil {
			fmt.Fprintf(os.Stderr, "error %d - %s: %s\n", i+1, repo.Name, errb.String())
			continue
		}
		if len(out) > 0 {
			if !repo.Auto_Commit {
				fmt.Printf("%d - %s: Uncommited changes\n\n", i+1, repo.Name)
				continue
			} else {
				addAndCommit := exec.Command("git", "add", ":", "&&", "git-commit", "--message=\"auto\"", "--dry-run")
				fmt.Println("command is", addAndCommit.String())
				addAndCommit.Dir = repo.Path
				addAndCommit.Stderr = &errb
				_, err := addAndCommit.Output()
				if err != nil {
					fmt.Fprintf(os.Stderr, "error %d - %s: %s\n", i+1, repo.Name, errb.String())
					continue
				}
				fmt.Fprintf(os.Stderr, "error %d - %s: %s\n", i+1, repo.Name, errb.String())
			}
		}

		// TODO: add --dry-run ??
		push := exec.Command("git", "push", "--all", "--dry-run")
		push.Dir = repo.Path
		push.Stderr = &errb

		_, err = push.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %d - %s: %s\n", i+1, repo.Name, errb.String())
			continue
		}

		fmt.Printf("%d - %s: %v\n", i+1, repo.Name, errb.String())

	}
}
