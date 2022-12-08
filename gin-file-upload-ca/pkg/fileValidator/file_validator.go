package fileValidator

import (
	"errors"
	"mime/multipart"
	"strconv"
	"strings"
)

// Note: This is just an example of validating File. You may modify this error mechanism based on your error handling style
func Validate(file *multipart.FileHeader, maxMBSize int, mimeType string) error {
	if file == nil {
		return errors.New("file is required")
	}

	if file.Size > int64(maxMBSize<<20) {
		return errors.New("the file size must less then " + strconv.Itoa(maxMBSize) + " MB")
	}

	if !strings.Contains(file.Header["Content-Type"][0], mimeType) {
		return errors.New("the file must be an " + mimeType + " type")

	}
	return nil
}
