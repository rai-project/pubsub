package redis

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/serializer"
	_ "github.com/rai-project/serializer/bson"
	_ "github.com/rai-project/serializer/json"
	"github.com/rai-project/utils"
	"github.com/rai-project/vipertags"
)

type redisConfig struct {
	Provider       string                `json:"provider" config:"pubsub.provider"`
	Endpoints      []string              `json:"endpoints" config:"pubsub.endpoints"`
	Password       string                `json:"password" config:"pubsub.password"`
	Serializer     serializer.Serializer `json:"-" config:"-"`
	SerializerName string                `json:"serializer_name" config:"broker.serializer" default:"json"`
	Cert           string                `json:"cert" config:"pubsub.cert"`
	done           chan struct{}         `json:"-" config:"-"`
}

var (
	Config = &redisConfig{
		done: make(chan struct{}),
	}
)

func (redisConfig) ConfigName() string {
	return "Redis"
}

func (a *redisConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *redisConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	a.Serializer, _ = serializer.FromName(a.SerializerName)
	if utils.IsEncryptedString(a.Password) {
		s, err := utils.DecryptStringBase64(config.App.Secret, a.Password)
		if err == nil {
			a.Password = s
		}
	}
}

func (c redisConfig) Wait() {
	<-c.done
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
