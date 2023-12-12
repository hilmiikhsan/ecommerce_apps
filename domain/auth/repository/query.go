package repository

const (
	queryCreate = `
	INSERT INTO auth (email, password,) VALUES (:email, :password)
	`

	queryGetByEmail = `
	SELECT
		id,
		email,
		password,
		role
	FROM auth
	WHERE email = $1
	`
	queryUpdateRole = `
	UPDATE auth SET role = $1 WHERE id = $2
	`
)
