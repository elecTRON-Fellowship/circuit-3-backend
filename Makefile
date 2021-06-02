postgres:
	sudo docker run --name f1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:13.3-alpine

createdb:
	sudo docker exec -it f1 createdb --username=root --owner=root formula_1

dropdb:
	sudo docker exec -it f1 dropdb formula_1

sqlc:
	sqlc generate

migrateup:
	migrate -path internal/auth/db/migration -database "postgresql://root:postgres@localhost:5432/formula_1?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/auth/db/migration -database "postgresql://root:postgres@localhost:5432/formula_1?sslmode=disable" -verbose down

test:
	go test -race -cover ./...

.PHONY: postgres createdb dropdb sqlc migrateup migratedown test
