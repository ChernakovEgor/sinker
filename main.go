package main

import (
	"fmt"
	"os"

	_ "github.com/pelletier/go-toml/v2"
	// toml "repo_pusher/pkg/toml_parser"
)

func main() {
	data, err := os.ReadFile("config.toml")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}
