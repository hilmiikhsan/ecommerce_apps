package entity

import (
	"errors"
	"time"

	"github.com/ecommerce/dto"
)

var (
	ErrNameIsRequired        = errors.New("name is required")
	ErrDateOfBirtIsRequired  = errors.New("date of birt is required")
	ErrPhoneNumberIsRequired = errors.New("phone number is required")
	ErrGenderIsInvalid       = errors.New("gender is invalid")
	ErrAddressIsRequired     = errors.New("address is required")
	ErrImageURLIsInvalid     = errors.New("image url is invalid")
	layout                   = "2006-01-02"
)

type User struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birt"`
	PhoneNumber string `db:"phone_number"`
	Gender      string `db:"gender"`
	Address     string `db:"address"`
	ImageURL    string `db:"image"`
	IsActive    bool   `db:"is_active"`
	CreatedBy   string `db:"created_by"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func NewUser() User {
	return User{}
}

func (u User) Validate(req dto.CreateOrUpdateUserRequest, id string) (User, error) {
	if req.Name == "" {
		return u, ErrNameIsRequired
	}

	if _, err := time.Parse(layout, req.DateOfBirth); err != nil {
		return u, ErrDateOfBirtIsRequired
	}

	if len(req.PhoneNumber) <= 10 {
		return u, ErrPhoneNumberIsRequired
	}

	if err := u.CheckUserGender(u.Gender); err != nil {
		return u, ErrGenderIsInvalid
	}

	if req.Address == "" {
		return u, ErrAddressIsRequired
	}

	if req.ImageURL == "" {
		return u, ErrImageURLIsInvalid
	}

	user := User{
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Gender:      req.Gender,
		Address:     req.Address,
		ImageURL:    req.ImageURL,
		CreatedBy:   id,
	}

	return user, nil
}

func (u User) CheckUserRole(role string) error {
	if role != "merchant" {
		return ErrInvalidRole
	}

	return nil
}

func (u User) CheckUserGender(gender string) error {
	switch {
	case gender == "male":
		return nil
	case gender == "female":
		return nil
	}

	return ErrGenderIsInvalid
}

func (u User) GetUserProfileResponse(user User) dto.GetUserProfileResponse {
	response := dto.GetUserProfileResponse{
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Address:     user.Address,
		ImageURL:    user.ImageURL,
	}

	return response
}
