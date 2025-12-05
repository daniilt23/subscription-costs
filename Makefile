run:
	docker compose up --build
swag:
	cd backend && swag init -g ./cmd/main.go
lint:
	cd backend && golangci-lint run