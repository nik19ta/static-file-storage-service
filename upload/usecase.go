package upload

import "mime/multipart"

type UseCase interface {
	Load(file *multipart.FileHeader) (string, error)
}
