package router

import (
	"CMS/controller"
	"CMS/dependency"
	"github.com/gin-gonic/gin"
)

func InitRouter() (*gin.Engine, error) {
	r := gin.New()

	api := r.Group("/api")
	initDocumentRouter(api)
	return r, nil
}

func initDocumentRouter(group *gin.RouterGroup) {
	documentController := dependency.Get("CMS.controller.DocumentController").(controller.DocumentController)

	documentGroup := group.Group("/document")
	documentGroup.POST("/", documentController.CreateDocument)

}
