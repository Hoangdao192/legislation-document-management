package controller

import (
	"CMS/dto"
	"CMS/model"
	"CMS/service"
	"encoding/json"
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
		return
	}

	var metadata map[string]interface{}
	if req.Metadata != "" {
		err = json.Unmarshal([]byte(req.Metadata), &metadata)
		if err != nil {
			res = MakeBadRequestResponse[*model.Document]("Cannot read provided metadata")
			ctx.JSON(400, MakeBadRequestResponse[*model.Document])
			return
		} else {
			req.ParsedMetadata = metadata
		}
	}

	res = c.documentService.CreateDocument(ctx, req)

	ctx.JSON(res.Code, res)
}
