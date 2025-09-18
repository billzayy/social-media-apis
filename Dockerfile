# Use Go base image for building
FROM golang:latest AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create working directory for api-gateway
WORKDIR /app/api-gateway

# Copy go.mod and go.sum files
COPY ./api-gateway/go.mod ./api-gateway/go.sum ./

# Copy authen-service folder for local module replacement
COPY ./authen-service ../authen-service
COPY ./chat-service ../chat-service
COPY ./post-service ../post-service
COPY ./user-service ../user-service
COPY ./notification-service ../notification-service

# Download dependencies
RUN go mod download

# Copy only the api-gateway service code
COPY ./api-gateway .

# Build the Go app
RUN go build -o api-gateway ./cmd/main.go

# Final stage
FROM alpine:3.20

# Set working directory
WORKDIR /root/

# Install certificates
RUN apk --no-cache add ca-certificates

# Create directory for .env file
RUN mkdir -p /root/internal

# Copy built binary from builder to a specific file path
COPY --from=builder /app/api-gateway/api-gateway /root/

# Copy .env file
COPY ./api-gateway/internal/.env /root/internal/.env

# Expose port defined by REST_PORT environment variable
EXPOSE $REST_PORT

# Run the app
CMD ["./api-gateway"]
# CMD ["sleep", "infinity"]