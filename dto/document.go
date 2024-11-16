package dto

import "mime/multipart"

type CreateDocumentRequest struct {
	Name           string                `form:"name" binding:"required"`
	File           *multipart.FileHeader `form:"file" binding:"required"`
	Metadata       string                `form:"metadata"`
	ParsedMetadata map[string]interface{}
}
