build:
	@go build -o bin/api
run: build
	@./bin/api
seeds:
	@go run ./scripts/seeds.go
test:
	@go test -v ./...
