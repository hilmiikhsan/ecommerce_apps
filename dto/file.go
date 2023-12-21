package dto

type UploadFileResponse struct {
	Url string `json:"url"`
}

func NewUploadFileResponse(url string) UploadFileResponse {
	return UploadFileResponse{
		Url: url,
	}
}
