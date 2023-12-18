test:
	go test -v ./... -cover

run:
	go run app/main.go serve-http

goose-create:
# example : make goose-create name=create_users_table

ifndef name
	$(error Usage: make goose-create name=<table_name>)
else
	@goose -dir pkg/database/migrations create $(name) sql
endif

goose-up:
# example : make goose-up
	@goose -dir pkg/database/migrations postgres "host=localhost port=39543 user=root-admin password=root-password dbname=ecommerce sslmode=disable" up

goose-down:
# example : make goose-down
	@goose -dir pkg/database/migrations postgres "host=localhost port=39543 user=root-admin password=root-password dbname=ecommerce sslmode=disable" down

goose-status:
# example : make goose-status
	@goose -dir pkg/database/migrations postgres "host=localhost port=39543 user=root-admin password=root-password dbname=ecommerce sslmode=disable" status