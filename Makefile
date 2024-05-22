build:
	@go build -o bin/service

run : build
	@./bin/service

watch:
	@reflex -r '\.go$$' -s -- sh -c 'make run'

dev:
	@go run main.go

test:
	@go test -v ./...