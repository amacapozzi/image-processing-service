package services

import (
	"github.com/h2non/filetype"
)

func IsValidImage(imgBytes []byte) bool {
	return filetype.IsImage(imgBytes)
}
