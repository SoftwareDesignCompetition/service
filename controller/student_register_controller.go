package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/service/service"
	log "github.com/cihub/seelog"
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
	err := s.StudentRegisterService.Register()
	if err!=nil {
		log.Errorf("studentRegister error %v",err)
		c.JSON(400, gin.H{
			"result":  1,
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