migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose down
.PHONY: migrateup