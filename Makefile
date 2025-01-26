build:
	@go build -o bin/ecom-golang-backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-golang-backend