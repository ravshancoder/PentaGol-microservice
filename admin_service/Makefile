run:
	go run cmd/main.go

proto-gen:
	./scripts/gen-proto.sh

migrate_up:
	migrate -path migrations -database postgres://ravshan:r@localhost:5432/admindb -verbose up

migrate_down:
	migrate -path migrations -database postgres://ravshan:r@localhost:5432/admindb -verbose down

migrate_force:
	migrate -path migrations -database postgres://ravshan:r@localhost:5432/admindb -verbose force 0

.PHONY: start migrateup migratedown