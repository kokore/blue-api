package service

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
)

type Services struct {
}

func Init(appConfig *config.AppConfig, repos repository.Repositories) Services {

	return Services{}
}
