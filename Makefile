.PHONY: postgres migrate

pg-up:
	nerdctl run -d --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v postgres_data:/var/lib/postgresql/data postgres
pg-down:
	nerdctl rm -f postgres

mgr-up:
	migrate -source file://migrations -database postgres://postgres:postgres@localhost/postgres?sslmode=disable up

mgr-down:
	migrate -source file://migrations -database postgres://postgres:postgres@localhost/postgres?sslmode=disable down
