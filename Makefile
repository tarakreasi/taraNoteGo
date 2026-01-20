.PHONY: run build clean migrate seed build-assets build-linux build-windows release dist

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

dist: release
	mkdir -p dist/linux
	mkdir -p dist/windows
	# Linux Bundle
	cp bin/taranote-linux dist/linux/server
	cp bin/migrate dist/linux/migrate
	cp -r views dist/linux/views
	cp -r public dist/linux/public
	mkdir -p dist/linux/database
	cp .env.example dist/linux/.env.example
	# Windows Bundle
	cp bin/taranote-windows.exe dist/windows/server.exe
	cp -r views dist/windows/views
	cp -r public dist/windows/public
	mkdir -p dist/windows/database
	cp .env.example dist/windows/.env.example
	@echo "Distribution created in dist/"

migrate:
	go run cmd/migrate/main.go

seed:
	go run cmd/seed/main.go

clean:
	rm -rf bin public/build dist
