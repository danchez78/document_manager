setup:
	go run github.com/swaggo/swag/cmd/swag@v1.16 \
		f -d "internal/common/server,internal/application/infrastructure/api"
	go run github.com/swaggo/swag/cmd/swag@v1.16 \
		i -g "server.go" -d "internal/common/server,internal/application/infrastructure/api" --output api
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run
