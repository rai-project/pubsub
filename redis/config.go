package redis

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type redisConfig struct {
	Provider  string   `json:"provider" config:"pubsub.provider"`
	Endpoints []string `json:"endpoints" config:"pubsub.endpoints"`
	Password  string   `json:"password" config:"pubsub.password"`
	Cert      string   `json:"cert" config:"pubsub.cert"`
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
