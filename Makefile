IMAGE_NAME = shortify-app
CONTAINER_NAME = shortify-container
BINARY_NAME = shortify

.PHONY: lint
lint:
	@echo "Generate swager docs"
	@swag init -g ./src/main.go
	@echo "Running golangci-lint with custom config..."
	@golangci-lint run ./...
	@if [ $$? -ne 0 ]; then \
		echo "Linting failed. Fix the issues before proceeding."; \
		exit 1; \
	fi

.PHONY: build
build: lint
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .

.PHONY: run
run: build
	@echo "Running Docker container..."
	docker run --name $(CONTAINER_NAME) -d -p 8080:8080 $(IMAGE_NAME)

.PHONY: dev
dev: lint
	@echo "Up with docker-compose"
	ENVIRONMENT=dev docker-compose up

.PHONY: sandbox
sandbox: lint
	@echo "Up with docker-compose"
	ENVIRONMENT=sandbox docker-compose up
