package controller

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/service/middleware"
	"github.com/service/service"
)

type StudentRegisterController struct {
	BaseController
	StudentRegisterService *service.StudentRegisterService
}

func NewStudentRegisterController(context *service.AppContext) (*StudentRegisterController, error) {
	studentRegisterController := context.Services["studentRegisterService"].(*service.StudentRegisterService)
	return &StudentRegisterController{
		BaseController: BaseController{
			Config: context.Config,
		},
		StudentRegisterService: studentRegisterController,
	}, nil
}
func (s *StudentRegisterController) Register(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "sig is empty",
		})
		return
	}
	param["signal"] = signal
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	name := c.PostForm("name")
	if name == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "name is empty",
		})
		return
	}
	param["name"] = name
	address := c.PostForm("address")
	if address == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "address  is empty",
		})
		return
	}
	param["address"] = address
	subject := c.PostForm("subject")
	if subject == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "subject  is empty",
		})
		return
	}
	param["subject"] = subject
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey  is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	param["appkey"] = appkey
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err := s.StudentRegisterService.Register(name, phone, subject, address)
	if err != nil {
		log.Errorf("studentRegister error %v", err)
		c.JSON(400, gin.H{
			"result":  2,
			"message": "studentRegister error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}
