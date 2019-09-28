package model

import (
	log "github.com/cihub/seelog"
	"github.com/service/config"
	"github.com/tb_common/cache"
)

type StudentRegisterModel struct {
	BaseCacheModel
}

func NewStudentRegisterModel(config *config.AppConfig) (*StudentRegisterModel, error) {
	caches, err := cache.GetCaches(&config.Redis)
	if err != nil {
		log.Errorf("get caches error: %v", err)
		return nil, err
	}
	return &StudentRegisterModel{
		BaseCacheModel: BaseCacheModel{
			Config: config,
			Dbs:    caches,
		},
	}, nil
}

func (s *StudentRegisterModel) Register() error {
	return nil
}
