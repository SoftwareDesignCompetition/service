package model

import (
	log "github.com/cihub/seelog"
	"github.com/service/config"
	"github.com/service/rediskey"
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

func (s *StudentRegisterModel) Register(name, phone, subject, address string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":name", name).Result()
	if err != nil {
		log.Errorf("add name err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":subject", subject).Result()
	if err != nil {
		log.Errorf("add subject err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":address", address).Result()
	if err != nil {
		log.Errorf("add address err, %v", err)
		return err
	}
	return nil
}
