generate_api:
	go run github.com/swaggo/swag/cmd/swag@v1.16 \
		f -d "internal/application/infrastructure/api/server,internal/application/infrastructure/api"
	go run github.com/swaggo/swag/cmd/swag@v1.16 \
		i -g "server.go" -d "internal/application/infrastructure/api/server,internal/application/infrastructure/api" --output api

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run
