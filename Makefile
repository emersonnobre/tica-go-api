run-swagger:
	cd src
	swag init --dir src --output src/docs

run-dev:
	@echo "Iniciando servidor"
	cd src
	swag init --dir src --output src/docs
	go run src/cmd/seed/main.go
	go run src/main.go development
run-prod:
	go run src/main.go production

run-container:
	docker-compose up -d --build

migration:
	migrate create -ext sql -dir src/cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate-dev-up:
	go run src/cmd/migrate/main.go development up $(filter-out $@,$(MAKECMDGOALS))
migrate-dev-down:
	go run src/cmd/migrate/main.go development down
migrate-prod-up:
	go run src/cmd/migrate/main.go production up
migrate-prod-down:
	go run src/cmd/migrate/main.go production down
seed-database:
	go run src/cmd/seed/main.go

generate-json:
	./src/scripts/create_json/run -file src/internal/core/domain/$(filter-out $@,$(MAKECMDGOALS))
generate-sql-table:
	./src/scripts/create_table/run -file src/internal/core/domain/$(filter-out $@,$(MAKECMDGOALS))