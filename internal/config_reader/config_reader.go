package configreader

import (
	"errors"
	"fmt"
	"os"

	"github.com/ChernakovEgor/sinker/pkg/assert"
	"github.com/pelletier/go-toml/v2"
)

type ServerConfig struct {
	Name   string
	Type   string
	Remote string
	Host   string
	Port   int
}

type RepoConfig struct {
	Name string
	Path string
}

type Config struct {
	Servers []ServerConfig
	Repos   []RepoConfig
}

func ReadConfig(path string) (*Config, error) {
	assert.Assert(path != "", "path should not be empty")

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ReadConfig: could not open file, %v", err))
	}
	defer f.Close()

	config := Config{}
	decoder := toml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ReadConfig: could not decode, %v", err))
	}

	return &config, nil
}
