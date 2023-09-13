package config

import (
	"blue-api/internal/database"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	APIConfig      APIConfig
	CORSConfig     CORSConfig
	DatabaseConfig database.DBConfig
}

type APIConfig struct {
	Env     string `envconfig:"ENV" default:"local"`
	BaseURL string `envconfig:"BASE_URL"`
	Port    string `envconfig:"PORT" default:"3000"`
}

func (a APIConfig) IsProd() bool {
	return a.Env == "production"
}

type LogConfig struct {
	Level string `envconfig:"LOG_CONFIG_LEVEL" default:"info"`
}

type CORSConfig struct {
	AllowSubDomainURL string `envconfig:"CORS_ALLOW_SUB_DOMAIN_URL"`
}

func (cfg *AppConfig) Init() {
	envconfig.MustProcess("", &cfg.APIConfig)
	envconfig.MustProcess("", &cfg.CORSConfig)
	envconfig.MustProcess("", &cfg.DatabaseConfig)
}

func LoadConfig() *AppConfig {
	_, ok := os.LookupEnv("ENV")
	if !ok {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		err := godotenv.Load(fmt.Sprintf("%v/../../.env.local", basePath))
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
	}

	cfg := &AppConfig{}
	cfg.Init()
	return cfg
}
