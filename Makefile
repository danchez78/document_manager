generate_api:
	go run github.com/swaggo/swag/cmd/swag@latest \
		f -d "internal/application/infrastructure/server"
	go run github.com/swaggo/swag/cmd/swag@latest \
		i -g "server.go" -d "internal/application/infrastructure/server" --output api
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
		run
