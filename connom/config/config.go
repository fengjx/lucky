package config

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/fengjx/go-halo/fskit"
	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/env"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"
)

var appConfig AppConfig

func init() {
	configArg := flag.String("c", "", "custom config file path")
	flag.Parse()
	appCfg := "conf/app.yml"
	if !env.IsDev() {
		exePath, err := os.Executable()
		if err != nil {
			log.Panic("can not get exec file path", zap.Error(err))
		}
		appCfg = filepath.Join(filepath.Dir(exePath), "conf/app.yml")
	}
	log.Infof("appCfg: %s", appCfg)
	appConfigFile, err := fskit.Lookup(appCfg, 5)
	if err != nil {
		log.Panic("config file not found", zap.String("path", appCfg), zap.Error(err))
	}
	configs := []string{appConfigFile}
	var configFile string
	if configArg != nil {
		configFile = *configArg
	}
	if configFile == "" {
		envConfig := os.Getenv("APP_CONFIG")
		if envConfig != "" {
			configFile = envConfig
		}
	}
	if configFile != "" {
		configFile, err = fskit.Lookup(configFile, 5)
		if err != nil {
			log.Panic("config file not found", zap.String("path", configFile), zap.Error(err))
		}
		log.Infof("load custom config file: %s", configFile)
		configs = append(configs, configFile)
	}
	appConfig = luchen.MustLoadConfig[AppConfig](configs...)
}

type AppConfig struct {
	Server Server                  `json:"server"`
	Auth   AuthConfig              `json:"auth"`
	DB     map[string]*DbConfig    `json:"db"`
	Redis  map[string]*RedisConfig `json:"redis"`
}

type Server struct {
	HTTP HTTPServerConfig
}

type AuthConfig struct {
	Version string `json:"version"`
	Secret  string `json:"secret"`
}

// CorsConfig 跨域配置
type CorsConfig struct {
	AllowOrigins []string `json:"allow-origins"`
}

type HTTPServerConfig struct {
	ServerName string     `json:"server-name"`
	Listen     string     `json:"listen"`
	Cors       CorsConfig `json:"cors"`
}

type DbConfig struct {
	Type    string `json:"type"`
	Dsn     string `json:"dsn"`
	MaxIdle int    `json:"max-idle"`
	MaxConn int    `json:"max-conn"`
	ShowSQL bool   `json:"show-sql"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func GetConfig() AppConfig {
	return appConfig
}
