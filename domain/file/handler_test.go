package file

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ecommerce/config"
	"github.com/ecommerce/infra/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

var handler = FileHandler{}

type mockFileService struct{}

// UploadFile implements Service.
func (mockFileService) UploadFile(ctx context.Context, buffer *bytes.Buffer, path string) (uri string, err error) {
	return UploadFileHandler()
}

var (
	UploadFileHandler func() (uri string, err error)
	jwtSecret         config.JWT
)

func init() {
	mock := mockFileService{}

	handler = NewFileHandler(mock)
}

func TestMain(m *testing.M) {
	err := config.LoadConfig("../../config/config.yaml")
	if err != nil {
		panic(err)
	}
	jwtSecret = config.Cfg.JWT
	middleware.SetJWTSecretKey(jwtSecret.Secret)
	m.Run()
}

func TestUploadFileHandler(t *testing.T) {
	type testCase struct {
		title              string
		expectedErr        error
		expectedStatusCode int
		requestFormValue   map[string]string
		endpoint           string
		contentType        string
		requestHeader      string
		before             func() error
	}

	var testCases = []testCase{
		{
			title:       "upload file success",
			expectedErr: nil,
			requestFormValue: map[string]string{
				"file": "file",
				"type": "type",
			},
			expectedStatusCode: fiber.StatusOK,
			endpoint:           "/v1/files/upload",
			contentType:        "multipart/form-data",
			requestHeader:      "Bearer ",
			before: func() error {
				UploadFileHandler = func() (uri string, err error) {
					return "https://cloudinary.com/ecommerce/1.png", nil
				}
				return nil
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			router := fiber.New()

			beforeErr := test.before()

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.Claims{
				ID:    "1",
				Email: "user@gmail.com",
				Role:  "user",
			})
			signedToken, err := token.SignedString([]byte(jwtSecret.Secret))
			require.NoError(t, err)

			mockService := mockFileService{}
			handler := NewFileHandler(mockService)

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", "testfile.png")
			require.NoError(t, err)

			sample, err := os.Open("testfile/testfile.png")
			require.NoError(t, err)

			_, err = io.Copy(part, sample)
			require.NoError(t, err)
			require.NoError(t, writer.Close())

			contentType := "multipart/form-data; boundary=" + writer.Boundary()

			router.Post("/v1/files/upload", middleware.AuthMiddleware(), handler.Upload)

			request := httptest.NewRequest(fiber.MethodPost, test.endpoint, body)
			request.Header.Set(fiber.HeaderContentType, contentType)
			request.Header.Set(fiber.HeaderAuthorization, "Bearer "+signedToken)

			resp, _ := router.Test(request, 1)

			require.Equal(t, test.expectedStatusCode, resp.StatusCode)
			require.Equal(t, test.expectedErr, beforeErr)
		})
	}
}
