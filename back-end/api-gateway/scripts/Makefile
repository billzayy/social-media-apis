# Variables
ENVIRONMENT_PATH=./internal/.env
DOCKER_COMPOSE_FILE=./build/docker-compose.yml
DOCKER_COMPOSE_CMD=docker-compose -f $(DOCKER_COMPOSE_FILE) --env-file $(ENVIRONMENT_PATH)

# Default target
all: swagger run-docker run-gateway

doc: swagger run-gateway

swagger:
	@echo "Generate Swagger files ..."
	swag fmt
	swag init -g ./cmd/main.go -o docs

run-gateway:
	@echo "Running API Gateway..."
	go run ./cmd/main.go

# Run the Docker container using Docker Compose
run-docker:
	@echo "Starting services with Docker Compose..."
	$(DOCKER_COMPOSE_CMD) up -d

stop:
	@echo "Stopping services with Docker Compose ..."
	$(DOCKER_COMPOSE_CMD) stop

# Stop and remove Docker containers
down:
	@echo "Stopping and removing Docker containers..."
	$(DOCKER_COMPOSE_CMD) down
