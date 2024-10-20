package service

import (
	"os"
	"strings"

	"github.com/fengjx/go-halo/errs"
	"github.com/fengjx/go-halo/fskit"
	"github.com/fengjx/luchen/log"

	"github.com/fengjx/lucky/integration/db"
)

var AppSvc = &appService{}

type appService struct {
}

func (svc appService) Init() error {
	initSQLFile, err := fskit.Lookup("conf/init.sql", 5)
	if err != nil {
		return errs.Wrap(err, "conf/init.sql not found")
	}
	content, err := os.ReadFile(initSQLFile)
	if err != nil {
		return errs.Wrap(err, "read conf/init.sql err")
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
			return errs.Wrapf(err, "exec conf/init.sql err: %s", script)
		}
	}
	return nil
}
