run-dev:
	cd src
	swag init --dir src --output src/docs
run-prod:
	go run src/main.go production

run-container:
	docker-compose up -d --build

migration:
	migrate create -ext sql -dir src/cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate-dev-up:
	go run src/cmd/migrate/main.go development up
migrate-dev-down:
	go run src/cmd/migrate/main.go development down
migrate-prod-up:
	go run src/cmd/migrate/main.go production up
migrate-prod-down:
	go run src/cmd/migrate/main.go production down

generate-json:
	./scripts/create_json/run -file internal/core/domain/$(filter-out $@,$(MAKECMDGOALS))
generate-sql-table:
	./scripts/create_table/run -file internal/core/domain/$(filter-out $@,$(MAKECMDGOALS))