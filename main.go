package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/service/config"
	"github.com/service/controller"
	"github.com/service/middleware"
	"github.com/service/model"
	"github.com/service/service"
	"github.com/service/task"
	"github.com/tb_common/health"
	mlog "github.com/tb_common/middleware"
	prom "github.com/tb_common/middleware"
	"github.com/tb_common/perf"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	version = ""
)

func main() {
	showVersion := flag.Bool("version", false, "Show version")
	configFlag := flag.String("config", "config/config.yml", "configuration file")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var err error
	var appConfig config.AppConfig
	if err = configor.Load(&appConfig, *configFlag); err != nil {
		log.Criticalf("load config error: %v", err)
		os.Exit(1)
	}

	var logger log.LoggerInterface
	logger, err = log.LoggerFromConfigAsFile(appConfig.Logger)
	if err != nil {
		log.Errorf("init logger from %s error: %v", appConfig.Logger, err)
	} else {
		log.ReplaceLogger(logger)
	}
	defer log.Flush()
	log.Infof("Started Application at %v", time.Now().Format("January 2, 2006 at 3:04pm (MST)"))
	log.Infof("Version: %v", version)
	marshal, _ := json.Marshal(appConfig)
	log.Infof("Config: %v", string(marshal))

	runtime.GOMAXPROCS(runtime.NumCPU())

	perf.Init(appConfig.PprofAddrs)
	if appConfig.Monitor {
		health.InitMonitor(appConfig.MonitorAddrs)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Global middleware
	router.Use(mlog.Logger())
	router.Use(gin.Recovery())

	p := prom.NewPrometheus("gin")
	p.Use(router)

	studentModel, err := model.NewStudentModel(&appConfig)
	if err != nil {
		log.Errorf("init studentModel error, %v", err)
	}

	appContext := &service.AppContext{
		Config:   &appConfig,
		Services: map[string]interface{}{},
		Models: map[string]interface{}{
			"studentModel": studentModel,
		},
	}
	appContext.Services["studentService"], _ = service.NewStudentService(appContext)

	authorized := router.Group("/v1/")
	authorized.Use(middleware.Auth())
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "login",
			})
		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "login",
			})
		})
	}
	studentController, _ := controller.NewStudentController(appContext)
	router.POST("/v1/student/Register", studentController.Register)
	router.POST("/v1/student/GetStudent", studentController.GetStudent)
	router.POST("/v1/student/ChangePhone", studentController.ChangePhone)
	router.POST("/v1/student/ChangeAddress", studentController.ChangeAddress)
	router.POST("/v1/student/ChangeGrade", studentController.ChangeGrade)
	router.POST("/v1/student/ChangeSubject", studentController.ChangeSubject)
	router.POST("/v1/student/ChangeName", studentController.ChangeName)

	taskManager, err := task.NewTaskManager(&appConfig, appContext)
	if err != nil {
		log.Errorf("new task basePlugin error, %v", err)
		return
	}
	taskManager.Start()

	srv := &http.Server{
		Addr:           appConfig.Addr, // listen and serve on 0.0.0.0:8080
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listen: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with  a timeout of 5 seconds.
	InitSignal()

	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown:", err)
	}
	log.Info("Server exist")
}
