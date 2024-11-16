package dto

import "mime/multipart"

type CreateDocumentRequest struct {
	Name     string                 `json:"name" binding:"required"`
	File     *multipart.FileHeader  `json:"file" binding:"required"`
	Metadata map[string]interface{} `json:"metadata"`
}
