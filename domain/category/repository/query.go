package repository

const (
	queryCreate = `
	INSERT INTO categories (name, created_by) VALUES (:name, :created_by)
	`

	queryGetAll = `
	SELECT
		id,
		name
	FROM categories
	`
)
