.PHONY: run build clean migrate seed build-assets build-linux build-windows release

run:
	npm run start

build-assets:
	npm run build

build-linux: build-assets
	GOOS=linux GOARCH=amd64 go build -o bin/taranote-linux cmd/server/main.go
	go build -o bin/migrate cmd/migrate/main.go

build-windows: build-assets
	GOOS=windows GOARCH=amd64 go build -o bin/taranote-windows.exe cmd/server/main.go

release: build-linux build-windows

migrate:
	go run cmd/migrate/main.go

seed:
	go run cmd/seed/main.go

clean:
	rm -rf bin public/build
