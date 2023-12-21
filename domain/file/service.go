package file

import (
	"bytes"
	"context"

	"github.com/ecommerce/infra/storage/images"
)

type Service interface {
	UploadFile(ctx context.Context, buffer *bytes.Buffer, path string) (uri string, err error)
}

type FileService struct {
	cloud images.CloudinaryService
}

func NewFileService(cloud images.CloudinaryService) FileService {
	return FileService{
		cloud: cloud,
	}
}

func (f FileService) UploadFile(ctx context.Context, buffer *bytes.Buffer, path string) (uri string, err error) {
	uri, err = f.cloud.Upload(ctx, buffer, path)
	if err != nil {
		return "", err
	}

	return uri, nil
}
