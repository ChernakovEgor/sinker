package main

import (
	"fmt"
	"sinker/internal/load"
)

func main() {
	fmt.Println(load.ReadConfig("blablabla"))
}
