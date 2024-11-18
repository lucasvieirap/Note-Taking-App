build:
	go build -o ./bin/main ./cmd/api/main.go
serve:
	go run ./cmd/api/main.go
