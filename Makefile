build:
	@go build -o bin/bin cmd/app/main.go

run: build
	@./bin/bin

test:
	@go test -v ./...

create-postgres-container:
	docker run --name postgres -e POSTGRES_PASSWORD=dbpass -p 5404:5432 -d postgres

start-postgres-container:
	docker start postgres

stop-postgres-container:
	docker stop postgres