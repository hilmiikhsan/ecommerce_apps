package images

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/ecommerce/config"
)

type CloudinaryService interface {
	Upload(ctx context.Context, file interface{}, pathDestination string) (uri string, err error)
}

type Cloudinary struct {
	client *cloudinary.Cloudinary
	cfg    config.FileCloudStorage
}

func NewCloudinary(client *cloudinary.Cloudinary, cfg config.FileCloudStorage) Cloudinary {
	return Cloudinary{
		client: client,
		cfg:    cfg,
	}
}

func CloudinaryStorage(cfg config.FileCloudStorage) (response Cloudinary, err error) {
	client, err := cloudinary.NewFromParams(cfg.CloudinaryName, cfg.CloudinaryAPIKey, cfg.CloudinaryAPISecret)
	if err != nil {
		return
	}

	return Cloudinary{
		client: client,
		cfg:    cfg,
	}, nil
}

func (c Cloudinary) Upload(ctx context.Context, file interface{}, pathDestination string) (uri string, err error) {
	filename := time.Now().Unix()
	res, err := c.client.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: "E-Commerce/" + pathDestination + "/" + fmt.Sprintf("%v", filename),
		Eager:    "q_10",
	})
	if err != nil {
		return "", err
	}

	if len(res.Eager) > 0 {
		return res.Eager[0].SecureURL, nil
	}

	url := res.SecureURL

	return url, nil
}
