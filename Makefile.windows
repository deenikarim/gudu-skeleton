BINARY_NAME=guduApp

build:
	@go mod vendor
	@echo "Building Gudu..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Gudu built!"

run: build
	@echo "Starting Gudu..."
	@./tmp/${BINARY_NAME} &
	@echo "Gudu started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker compose up -d

stop_compose:
	docker compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

s: run

stop:
	@echo "Stopping Gudu..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Gudu!"

restart: stop start