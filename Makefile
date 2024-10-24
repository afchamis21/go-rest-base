start:
	@go run cmd/api/main.go
	
run: build
	@./bin/

docker-db:
	@docker-compose up -d

migrate:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down