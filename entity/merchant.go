package entity

type Merchant struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Address     string `db:"address"`
	ImageUrl    string `db:"image_url"`
	City        string `db:"city"`
	Role        string `db:"role"`
	CreatedBy   string `db:"created_by"`
}
