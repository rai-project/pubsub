package redis

import (
	"strings"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/serializer"
	"github.com/rai-project/serializer/bson"
	"github.com/rai-project/serializer/json"
	"github.com/rai-project/vipertags"
)

type redisConfig struct {
	Provider       string                `json:"provider" config:"pubsub.provider"`
	Endpoints      []string              `json:"endpoints" config:"pubsub.endpoints"`
	Password       string                `json:"password" config:"pubsub.password"`
	Serializer     serializer.Serializer `json:"-" config:"-"`
	SerializerName string                `json:"serializer_name" config:"broker.serializer" default:"json"`
	Cert           string                `json:"cert" config:"pubsub.cert"`
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
	switch strings.ToLower(a.SerializerName) {
	case "json":
		a.Serializer = json.New()
	case "bson":
		a.Serializer = bson.New()
	default:
		log.WithField("serializer", a.SerializerName).
			Warn("Cannot find serializer")
	}
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
