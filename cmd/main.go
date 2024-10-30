package main

import (
	"fmt"
	"github.com/ChernakovEgor/sinker/internal/load"
)

func main() {
	fmt.Println(load.ReadConfig("/Users/egor/developer/pet/repo_pusher/config.toml"))
}
