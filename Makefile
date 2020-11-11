db: migrate
.PHONY: db

postgres:
	docker-compose -f docker-compose.yaml up -d db
	sleep 5

createdb: postgres
	docker exec -it db createdb --username=postgres --owner=postgres wastedb

migrate: createdb
	migrate -database postgres://postgres:postgres@localhost:15432/wastedb?sslmode=disable -path migrations up

app: db
	docker-compose -f docker-compose.yaml up -d app

start: app

stop:
	docker-compose -f docker-compose.yaml down

