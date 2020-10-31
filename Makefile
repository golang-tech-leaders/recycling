db: migrate
.PHONY: db

postgres:
	docker-compose -f docker-compose.yaml up -d
	sleep 2

createdb: postgres
	docker exec -it db createdb --username=postgres --owner=postgres wastedb

migrate: createdb
	migrate -database postgres://postgres:postgres@localhost:15432/wastedb?sslmode=disable -path migrations up

stop:
	docker-compose -f docker-compose.yaml down
