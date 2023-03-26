package usecase

import (
	"mime/multipart"
	"path/filepath"
	"static/upload"

	uuid "github.com/google/uuid"

	config "static/pkg/config"
)

type heavyApiLoad struct {
	heavyApiLoadRepository upload.HeavyApiLoadRepository
}

func NewHeavyApiLoadUseCase(heavyApiLoadRepository upload.HeavyApiLoadRepository) *heavyApiLoad {
	return &heavyApiLoad{heavyApiLoadRepository: heavyApiLoadRepository}
}

func (h *heavyApiLoad) Load(file *multipart.FileHeader) (string, error) {
	conf := config.GetConfig()

	name := uuid.New().String() + filepath.Ext(file.Filename)
	path := conf.Dir + name

	err := h.heavyApiLoadRepository.LoadFile(path, file)

	return name, err
}
