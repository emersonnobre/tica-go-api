run-dev:
	go run cmd/api/main.go development
run-prod:
	go run cmd/api/main.go production

create-dev-database:
	docker-compose up -d --build

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate-dev-up:
	go run cmd/migrate/main.go development up
migrate-dev-down:
	go run cmd/migrate/main.go development down
migrate-prod-up:
	go run cmd/migrate/main.go production up
migrate-prod-down:
	go run cmd/migrate/main.go production down

generate-json:
	./scripts/create_json/run -file internal/core/domain/$(filter-out $@,$(MAKECMDGOALS))
generate-sql-table:
	./scripts/create_table/run -file internal/core/domain/$(filter-out $@,$(MAKECMDGOALS))