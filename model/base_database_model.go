package model

import (
	"github.com/service/config"
	"github.com/tb_common/db"
)

type BaseDatabaseModel struct {
	Config *config.AppConfig
	DBs    *db.Databases
}

func (d *BaseDatabaseModel) GetDbMaster() *db.Database {
	return d.DBs.GetDatabase("master")
}

func (d *BaseDatabaseModel) GetLockDbMaster() *db.Database {
	return d.DBs.GetDatabase("lock_master")
}
