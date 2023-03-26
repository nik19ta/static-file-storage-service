package upload

import "mime/multipart"

type HeavyApiLoadRepository interface {
	LoadFile(fileName string, file *multipart.FileHeader) error
}
