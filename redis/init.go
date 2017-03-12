package redis

import (
	llog "log"
	"os"

	"github.com/Sirupsen/logrus"
	redis "gopkg.in/redis.v5"

	"github.com/rai-project/config"
	logger "github.com/rai-project/logger"
)

var (
	log *logrus.Entry
)

func init() {
	config.AfterInit(func() {
		log = logger.WithField("pkg", "pubsub/redis")

		if config.IsDebug {
			lg := llog.New(os.Stdout, "logger: ", llog.LstdFlags)
			redis.SetLogger(lg)
		}
	})
}
