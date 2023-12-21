package file

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc = FileService{}

type mockCloudinaryService struct{}

// Upload implements images.CloudinaryService.
func (mockCloudinaryService) Upload(ctx context.Context, file interface{}, pathDestination string) (uri string, err error) {
	return UploadFile()
}

var (
	UploadFile func() (uri string, err error)
)

func init() {
	mock := mockCloudinaryService{}

	svc = NewFileService(mock)
}

func TestUploadFile(t *testing.T) {
	type testCase struct {
		title         string
		expectedErr   error
		expectedValue string
		before        func()
	}

	var testCases = []testCase{
		{
			title:         "upload file success",
			expectedErr:   nil,
			expectedValue: "https://cloudinary.com/ecommerce/1.png",
			before: func() {
				UploadFile = func() (uri string, err error) {
					return "https://cloudinary.com/ecommerce/1.png", nil
				}
			},
		},
		{
			title:         "upload file failed internal server error",
			expectedErr:   errors.New("internal server error"),
			expectedValue: "",
			before: func() {
				UploadFile = func() (uri string, err error) {
					return "", errors.New("internal server error")
				}
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			test.before()

			buffer := bytes.NewBufferString("test file content")

			uri, err := svc.UploadFile(context.Background(), buffer, "ecommerce")
			require.Equal(t, test.expectedErr, err)
			require.Equal(t, test.expectedValue, uri)
		})
	}
}
