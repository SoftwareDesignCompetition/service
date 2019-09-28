package config
import(
	cp "github.com/tb_common/cache"
	dbp "github.com/tb_common/db"
)
type AppConfig struct {
	Monitor                           bool     `default:"true"`							//程序状态？
	Logger                            string `default:"conf/logger.xml"`					//日志
	MonitorAddrs                      []string `default:"localhost:0816"`					//
	PprofAddrs                        []string `default:"localhost:0815"`					//监督网址
	Addr                              string   `default:"0.0.0.0:8080"`
	DB                                dbp.DBConfig
	Redis                             cp.RedisConfig
}