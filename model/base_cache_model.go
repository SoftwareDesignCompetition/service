package model

import (
	"github.com/service/config"
	"github.com/tb_common/cache"
)

type BaseCacheModel struct {
	Config *config.AppConfig
	Dbs    *cache.Caches
}

func (c *BaseCacheModel) GetRedisMaster() *cache.Cache {
	return c.Dbs.GetCache("master")
}
