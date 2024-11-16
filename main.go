package main

import (
	"CMS/config"
	"CMS/controller"
	"CMS/dependency"
	"CMS/model"
	"CMS/repository"
	"CMS/router"
	"CMS/service"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	initStorage()
	initRepository()
	initService()
	initController()

	r, err := router.InitRouter()
	if err != nil {
		panic(err)
	}
	_ = r.Run(fmt.Sprintf(":%d", config.ApplicationConfig.Port))
}

func initStorage() {
	path := config.ApplicationConfig.StorageDirectory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func initRepository() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ApplicationConfig.Mysql.User, config.ApplicationConfig.Mysql.Password,
		config.ApplicationConfig.Mysql.Host, config.ApplicationConfig.Mysql.Port,
		config.ApplicationConfig.Mysql.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Document{}, &model.File{})
	if err != nil {
		return
	}
	fileRepository := repository.NewFileRepository(db)
	documentRepository := repository.NewDocumentRepository(db)
	dependency.Add("CMS.repository.FileRepository", fileRepository)
	dependency.Add("CMS.repository.DocumentRepository", documentRepository)
}

func initService() {
	fileService := service.NewFileService(
		dependency.Get("CMS.repository.FileRepository").(repository.FileRepository))
	documentService := service.NewDocumentService(
		dependency.Get("CMS.repository.DocumentRepository").(repository.DocumentRepository),
		dependency.Get("CMS.repository.FileRepository").(repository.FileRepository),
		fileService,
	)
	dependency.Add("CMS.service.FileService", fileService)
	dependency.Add("CMS.service.DocumentService", documentService)
}

func initController() {
	documentController := controller.NewDocumentController(
		dependency.Get("CMS.service.DocumentService").(service.DocumentService),
	)
	dependency.Add("CMS.controller.DocumentController", documentController)
}
