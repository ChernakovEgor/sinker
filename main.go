package main

import (
	"fmt"
	toml "repo_pusher/pkg/toml_parser"
)

func main() {
	res := toml.ValidateTOML("")

	fmt.Println(res)
}
