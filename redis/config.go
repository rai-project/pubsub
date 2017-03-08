package redis

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type redisConfig struct {
	Provider     string   `json:"provider" config:"database.provider" default:"redis"`
	Endpoints    []string `json:"endpoints" config:"database.endpoints"`
	Username     string   `json:"username" config:"database.username"`
	Password     string   `json:"password" config:"database.password"`
	DatabaseName string   `json:"database_name" config:"database.database_name"`
}

var (
	Config = &redisConfig{}
)

func (redisConfig) ConfigName() string {
	return "Redis"
}

func (redisConfig) SetDefaults() {
}

func (a *redisConfig) Read() {
	vipertags.Fill(a)
}

func (c redisConfig) String() string {
	return pp.Sprintln(c)
}

func (c redisConfig) Debug() {
	log.Debug("Redis Config = ", c)
}

func init() {
	config.Register(Config)
}
