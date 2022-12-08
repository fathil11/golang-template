package repository

import (
	"io"
	"mime/multipart"
	"os"
)

const (
	// Note: should be set inside the ENV
	UPLOAD_DIRECTORY = "resources/upload/"
)

func Store(file *multipart.FileHeader, directory string, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Check or Create Directory
	_, err = os.Stat(UPLOAD_DIRECTORY + directory)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(UPLOAD_DIRECTORY+directory, os.ModePerm)
		} else {
			return err
		}
	}

	out, err := os.OpenFile(UPLOAD_DIRECTORY+directory+"/"+fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
