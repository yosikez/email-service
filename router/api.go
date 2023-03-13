package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yosikez/email-notification-services/controller"
)

func RegisterRouter(router *gin.Engine){
	mailController := controller.NewMailController()

	router.GET("/mails", mailController.FindAll)
	router.GET("/mails/:id", mailController.FindById)
	router.GET("/mails/category/:action", mailController.FindByAction)
}