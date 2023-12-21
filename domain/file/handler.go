package file

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/ecommerce/dto"
	logs "github.com/ecommerce/infra/logger"
	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
	service Service
}

func NewFileHandler(service Service) FileHandler {
	return FileHandler{
		service: service,
	}
}

func (f FileHandler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, ErrInvalidFileType)
	}

	// if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" && file.Header.Get("Content-Type") != "image/webp" {
	// 	err = ErrInvalidFileType
	// 	logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %v%v%v", err.Error(), "file type :", file.Header.Get("Content-Type")))
	// 	return WriteError(c, err)
	// }
	fileExt := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExts[strings.ToLower(fileExt)] {
		err = ErrInvalidFileType
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %v%v%v", err.Error(), "file type :", fileExt))
		return WriteError(c, err)
	}

	if file.Size > 1*1024*1024 {
		err = ErrInvalidFileSize
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %v%v%v%v", err.Error(), "file size :", (file.Size/1024/1024), "MB"))
		return WriteError(c, ErrInvalidFileSize)
	}

	typeFile := c.FormValue("type", "")

	source, err := file.Open()
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}
	defer source.Close()

	buffer := bytes.NewBuffer(nil)

	_, err = io.Copy(buffer, source)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	url, err := f.service.UploadFile(c.UserContext(), buffer, "ecommerce/"+typeFile)
	if err != nil {
		logs.Logger(logs.GetFunctionPath(), logs.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return WriteError(c, err)
	}

	payload := dto.NewUploadFileResponse(url)

	return WriteSuccess(c, "upload file success", payload, fiber.StatusOK)
}
