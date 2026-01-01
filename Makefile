.PHONY: run build clean migrate

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/migrate cmd/migrate/main.go

migrate:
	go run cmd/migrate/main.go

clean:
	rm -rf bin
