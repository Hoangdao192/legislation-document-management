package controller

import (
	"CMS/dto"
	"CMS/model"
	"CMS/service"
	"github.com/gin-gonic/gin"
)

type DocumentController interface {
	CreateDocument(*gin.Context)
}

type documentController struct {
	documentService service.DocumentService
}

func NewDocumentController(documentService service.DocumentService) DocumentController {
	return &documentController{
		documentService: documentService,
	}
}

func (c *documentController) CreateDocument(ctx *gin.Context) {
	var req dto.CreateDocumentRequest
	err := ctx.Bind(&req)
	var res *dto.BaseResponse[*model.Document]
	if err != nil {
		res = MakeBadRequestResponse[*model.Document](err.Error())
		ctx.JSON(400, MakeBadRequestResponse[string](err.Error()))
	}
	res = c.documentService.CreateDocument(ctx, req)

	ctx.JSON(res.Code, res)
}
