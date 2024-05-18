package service

import (
	"context"
	"os"
	"strings"

	"github.com/fengjx/go-halo/fskit"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/integration/db"
)

var AppSvc = &appService{}

type appService struct {
}

func (svc appService) Init(ctx context.Context) error {
	initSQLFile, err := fskit.Lookup("conf/init.sql", 5)
	if err != nil {
		log.ErrorCtx(ctx, "conf/init.sql not found")
		return err
	}
	content, err := os.ReadFile(initSQLFile)
	if err != nil {
		log.Error("read conf/init.sql err", zap.Error(err))
		return err
	}
	sqlScript := string(content)
	scripts := strings.Split(sqlScript, ";")
	for _, script := range scripts {
		if strings.TrimSpace(script) == "" {
			continue
		}
		log.Infof("ecec sql: %s", script)
		_, err = db.GetDefaultDB().Exec(script)
		if err != nil {
			log.Error("exec conf/init.sql err", zap.String("script", script), zap.Error(err))
			return err
		}
	}
	return nil
}
