run:
	@go run cmd/main.go

up: 
	@goose -dir=db/migrations sqlite db/app.db up

down: 
	@goose -dir=db/migrations sqlite db/app.db down