package usecase

import (
	"errors"
	"fathil/gin-file-upload-ca/pkg/fileValidator"
	"fathil/gin-file-upload-ca/repository"
	"fathil/gin-file-upload-ca/request"
	"strings"
	"time"
)

func Store(request *request.Request) error {
	var directory string

	switch request.Condition {
	case "customer logo":
		err := fileValidator.Validate(request.File, 2, "image")
		if err != nil {
			return err
		}

		directory = "customer"

	default:
		return errors.New("invalid request condition")
	}

	fileName := time.Now().Format("2006-01-02_150405_") + strings.ReplaceAll(request.File.Filename, " ", "_")
	err := repository.Store(request.File, directory, fileName)

	// Note: Call the database repository to save the fileName in the database

	return err
}
