package task

import (
	log "github.com/cihub/seelog"
	"github.com/service/config"
	"github.com/service/model"
	"github.com/service/service"
	"github.com/tb_common/cache"
	"github.com/tb_common/db"
)

type Manager struct {
	*model.BaseCacheModel
	*model.BaseDatabaseModel

	AppContext *service.AppContext
	Config     *config.AppConfig
}

func NewTaskManager(config *config.AppConfig, appContext *service.AppContext) (*Manager, error) {
	manager := &Manager{
		AppContext: appContext,
		Config:     config,
	}
	return manager, nil
}

func (s *Manager) init() {
	caches, err := cache.GetCaches(&s.Config.Redis)
	if err != nil {
		log.Errorf("get caches error: %v", err)
		return
	}
	s.BaseCacheModel = &model.BaseCacheModel{
		Config: s.Config,
		Dbs:    caches,
	}

	databases, err := db.GetDatabases(&s.Config.DB)
	if err != nil {
		log.Errorf("get databases error: %v", err)
		return
	}
	s.BaseDatabaseModel = &model.BaseDatabaseModel{
		Config: s.Config,
		DBs:    databases,
	}
}

//启动定时任务管理器
func (s *Manager) Start() {
	s.init()
}
