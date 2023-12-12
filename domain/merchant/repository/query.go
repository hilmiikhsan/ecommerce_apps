package repository

const (
	queryGetByCreatedBy = `
	SELECT
		m.id,
		m.name,
		m.phone_number,
		m.address,
		m.image_url,
		m.city,
		m.created_by,
		a.role
	FROM merchants m
	JOIN auth a ON a.id = m.created_by
	WHERE m.created_by = $1
	`
)
