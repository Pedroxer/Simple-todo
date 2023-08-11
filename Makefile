migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose down
test:
	go test -v ./go ...
server:
	go run main.go
.PHONY: migrateup, migratedown, test, server