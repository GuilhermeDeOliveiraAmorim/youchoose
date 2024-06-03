package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"youchoose/configs"

	"github.com/google/uuid"
)

func MoveFile(file multipart.File, handler *multipart.FileHeader) (int64, string, string, int64, error) {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	path := configs.LocalImagePath

	extension := filepath.Ext(handler.Filename)

	name := uuid.New().String()

	size := handler.Size

	fileCreate, err := os.Create(path + name + extension)
	if err != nil {
		return 0, "", "", 0, errors.New(err.Error())
	}

	defer file.Close()
	defer fileCreate.Close()

	fileWritten, err := io.Copy(fileCreate, file)
	if err != nil {
		return 0, "", "", 0, errors.New(err.Error())
	}

	extension = strings.Replace(filepath.Ext(handler.Filename), ".", "", -1)

	return fileWritten, name, extension, size, nil
}
