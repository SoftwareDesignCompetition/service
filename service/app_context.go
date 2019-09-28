package service

import "github.com/service/config"

type AppContext struct {
	Config *config.AppConfig

	Services map[string]interface{}
	Models   map[string]interface{}
}
