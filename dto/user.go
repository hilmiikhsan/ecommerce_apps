package dto

type CreateOrUpdateUserRequest struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageURL    string `json:"image_url"`
}
type GetUserProfileResponse struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	ImageURL    string `json:"image_url"`
}
