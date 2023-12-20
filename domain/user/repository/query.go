package repository

const (
	queryCreate = `INSERT INTO (name, date_of_birth, phone_number, gender, address, image_url) VALUES (:name, :date_of_birth, :phone_number, :gender, :address, :image_url)`

	queryGetProfile = `
	SELECT
		name,
		date_of_birth,
		phone_number,
		gender,
		address,
		image_url
	FROM users
	WHERE id = $1
	`

	queryUpdateProfile = `
	UPDATE users
	SET
		name = $2,
		date_of_birth = $3,
		phone_number = $4,
		gender = $5,
		address = $6,
		image_url = $7
	WHERE id = $1
	`
)
