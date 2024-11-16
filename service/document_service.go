package service

import (
	"CMS/dto"
	"CMS/model"
	"CMS/repository"
	"github.com/gin-gonic/gin"
)

type DocumentService interface {
	CreateDocument(*gin.Context, dto.CreateDocumentRequest) *dto.BaseResponse[*model.Document]
}

type documentService struct {
	documentRepository repository.DocumentRepository
	fileRepository     repository.FileRepository
	fileService        FileService
}

func NewDocumentService(
	documentRepository repository.DocumentRepository,
	fileRepository repository.FileRepository,
	fileService FileService) DocumentService {
	return &documentService{
		documentRepository: documentRepository,
		fileRepository:     fileRepository,
		fileService:        fileService,
	}
}

func (s *documentService) CreateDocument(ctx *gin.Context, req dto.CreateDocumentRequest) *dto.BaseResponse[*model.Document] {
	appFile := s.fileService.CreateFile(ctx, req.File)
	if appFile.Code != 200 {
		return MakeBadRequestResponse[*model.Document](appFile.Message)
	}

	document := model.Document{
		Name:           req.Name,
		SourceFileId:   appFile.Data.ID,
		PreviewFileId:  appFile.Data.ID,
		EditableFileId: appFile.Data.ID,
		Metadata:       req.ParsedMetadata,
		// TODO: Assign created by
	}
	err := s.documentRepository.Save(ctx, &document)
	if err != nil {
		return MakeBadRequestResponse[*model.Document]("Cannot save document")
	}

	return MakeSuccessResponse[*model.Document](&document)
}
