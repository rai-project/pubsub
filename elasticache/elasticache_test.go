package elasticache

import (
	"os"
	"testing"

	"github.com/rai-project/config"
)

func TestMain(m *testing.M) {
	config.Init(
		config.VerboseMode(true),
		config.DebugMode(true),
		config.ColorMode(true),
	)
	os.Exit(m.Run())
}
