package integration

import (
	"github.com/fengjx/luchen/log"

	_ "github.com/fengjx/lucky/integration/db"
)

func Init() {
	log.Info("integration init")
}
