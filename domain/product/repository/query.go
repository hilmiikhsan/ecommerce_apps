package repository

const (
	queryCreate = `
	INSERT INTO products (
		name, 
		description, 
		price, 
		stock, 
		category_id, 
		merchant_id, 
		image_url, 
		sku,
		created_by
	) VALUES (:name, :description, :price, :stock, :category_id, :merchant_id, :image_url, :sku, :created_by)
	`

	queryGetByMerchantId = `
	SELECT
		p.id,
		p.sku,
		p.name,
		p.description,
		p.price,
		p.stock,
		c.name as category,
		p.image_url
	FROM products p
	JOIN categories c ON c.id = p.category_id
	WHERE p.merchant_id = $1
	`

	queryCountByMerchantId = `
	SELECT COUNT(p.id) as total_data 
	FROM products p
	JOIN categories c ON c.id = p.category_id
	WHERE p.merchant_id = $1
	`

	queryGetById = `
	SELECT
		p.id,
		p.sku,
		p.name,
		p.description,
		p.price,
		p.stock,
		c.name as category,
		p.category_id,
		p.image_url,
		p.created_at,
		p.updated_at
	FROM products p
	JOIN categories c ON c.id = p.category_id
	WHERE p.id = $1
	`

	queryUpdate = `
	UPDATE products SET 
		name = :name, 
		description = :description, 
		price = :price, 
		stock = :stock, 
		category_id = :category_id, 
		image_url = :image_url, 
		updated_at = NOW() 
	WHERE id = :id
	`

	queryGetBySku = `
	SELECT
		p.id,
		p.sku,
		p.name,
		p.description,
		p.price,
		p.stock,
		c.name as category,
		p.category_id,
		p.merchant_id,
		p.image_url,
		p.created_at,
		p.updated_at,
		m.name as merchant_name,
		m.city as merchant_city
	FROM products p
	JOIN categories c ON c.id = p.category_id
	JOIN merchants m ON m.id = p.merchant_id
	WHERE p.sku = $1
	`
)
