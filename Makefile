# Variables
DOCKER_GO_IMAGE_NAME=coderbillzay/social-api-gateway
DOCKERFILE_GO_PATH=./Dockerfile
DOCKER_CONTEXT=.
DOCKER_COMPOSE_FILE=./docker-compose.yml
DOCKER_COMPOSE_CMD=docker-compose -f $(DOCKER_COMPOSE_FILE) 
ENV_FILE_PATH =./api-gateway/internal/.env

# Default target
all: swagger run-docker

# Format & Update swagger docs folder
swagger:
	@echo "Generate Swagger files ..."
	swag fmt -d ./api-gateway/cmd/main.go 
	swag init -d ./api-gateway/cmd -o docs

# Run the Docker container using Docker Compose
run-docker:
	@echo "Starting services with Docker Compose..."
	$(DOCKER_COMPOSE_CMD) up -d

# Run the Docker container using Docker Compose
run-docker-with-env:
	@echo "Starting services with Docker Compose..."
	$(DOCKER_COMPOSE_CMD) --env-file $(ENV_FILE_PATH) up -d

# Build the Docker image
build-image:
	@echo "Building Docker Go image $(DOCKER_GO_IMAGE_NAME)..."
	docker build -t $(DOCKER_GO_IMAGE_NAME) -f $(DOCKERFILE_GO_PATH) $(DOCKER_CONTEXT)

# Push image into DockerHub
push:
	@echo "Pushing into DockerHub ..."
	docker push $(DOCKER_GO_IMAGE_NAME)

# Clean cache and remove image
clean:
	@echo "Cleaning Docker images and build artifacts..."
	docker rmi $(DOCKER_GO_IMAGE_NAME)
	docker system prune --volumes -f%

# Stop Docker containers
stop:
	@echo "Stopping services with Docker Compose ..."
	$(DOCKER_COMPOSE_CMD) stop

# Stop and remove Docker containers
down:
	@echo "Stopping and removing Docker containers..."
	$(DOCKER_COMPOSE_CMD) down
