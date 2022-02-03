run:
	@echo "Running the application"
	go run cmd/main.go

test:
	@echo "Testing the application"
	go test -v ./...