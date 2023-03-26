package file

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type HeavyApiLoadRepository struct{}

func NewHeavyApiLoadRepository() *HeavyApiLoadRepository {
	return &HeavyApiLoadRepository{}
}

func (r *HeavyApiLoadRepository) LoadFile(name string, file *multipart.FileHeader) error {
	c := gin.Context{}

	err := c.SaveUploadedFile(file, name)

	return err
}
