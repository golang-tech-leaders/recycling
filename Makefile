#db: migrate
db: createdb
.PHONY: db

postgres:
	docker-compose up -d db
	until docker exec db pg_isready ; do sleep 5 ; done	

createdb: postgres
	docker exec -it db createdb --username=postgres --owner=postgres wastedb

#migrate: createdb
#	migrate -database postgres://postgres:postgres@localhost:15432/wastedb?sslmode=disable -path migrations up

app: db
	docker-compose up -d app

start: app

stop:
	docker-compose down

