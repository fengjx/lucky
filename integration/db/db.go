package db

import (
	"strings"

	"github.com/fengjx/daox"
	"github.com/fengjx/luchen/log"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	"github.com/fengjx/lucky/connom/config"
)

var (
	defaultDBName = "default"
	toLowerMapper = reflectx.NewMapperFunc("json", strings.ToLower)
	dbMap         = make(map[string]*sqlx.DB)
	txManagerMap  = make(map[string]*daox.TxManager)
)

func init() {
	for k, c := range config.GetConfig().DB {
		db, err := sqlx.Open(c.Type, c.Dsn)
		if err != nil {
			log.Panicf("create db connection err - %s, %s, %s", c.Type, c.Dsn, err.Error())
			return
		}
		err = db.Ping()
		if err != nil {
			log.Panicf("db ping err - %s, %s, %s", c.Type, c.Dsn, err.Error())
		}
		if c.MaxIdle != 0 {
			db.SetMaxIdleConns(c.MaxIdle)
		}
		if c.MaxConn != 0 {
			db.SetMaxOpenConns(c.MaxConn)
		}
		db.Mapper = toLowerMapper
		dbMap[k] = db
		txManagerMap[k] = daox.NewTxManager(db)
		log.Infof("init db[%s]", k)
	}
	defaultDB := dbMap[defaultDBName]
	daox.UseDefaultMasterDB(defaultDB)
	// 默认忽略字段，这两个字段交给mysql自行处理，主要是为了一些问题排查提供依据，即使手动修改数据库也会触发字段更新
	// 如果业务代码需要用到创建时间和更新时间，可以自行维护一个字段
	daox.UseSaveOmits("ctime", "utime")
}

func GetDefaultDB() *sqlx.DB {
	return dbMap[defaultDBName]
}

func GetDB(name string) *sqlx.DB {
	return dbMap[name]
}

func GetDefaultTxManager() *daox.TxManager {
	return txManagerMap[defaultDBName]
}

func GetTxManager(name string) *daox.TxManager {
	return txManagerMap[name]
}
