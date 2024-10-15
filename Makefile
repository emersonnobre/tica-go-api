run:
	@go run cmd/api/main.go
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate:
	@go run cmd/migrate/main.go