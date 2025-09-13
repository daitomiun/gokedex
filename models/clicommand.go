package models

import "github.com/daitomiun/gokedex/internal/pokecache"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, cache *pokecache.Cache, param string) error
}
