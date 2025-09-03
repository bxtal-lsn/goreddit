.PHONY: postgres migrate

postgres:
	docker run --rm -d -ti --network host -e POSTGRES_PASSWORD=secret postgres

migrate:
	migrate -source file://migrations -database postgres://postgres:secret@localhost/postgres?sslmode=disable up
