# Variables
DOCKER_GO_IMAGE_NAME=coderbillzay/social-user-service
DOCKERFILE_GO_PATH=./build/Dockerfile
DOCKER_CONTEXT=.

# Build command will run as clone
.PHONY: build

# Default target
all: allow generate build

docker: build push

install:
	go mod download

allow:
	chmod +x ./scripts/generate.sh 

# Run generate.sh script
generate:
	@echo "Running generate.sh..."
	./scripts/generate.sh

# Build the Docker image
build:
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
	docker system prune --volumes -f