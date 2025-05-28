# Use Go base image
FROM golang:latest AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy entire project
COPY . .

# Build the Go app
RUN go build -o user-service ./cmd/main.go

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Install certificates
RUN apk --no-cache add ca-certificates

# Copy built binary from builder
COPY --from=builder /app/user-service .

# Copy .env file
COPY ./internal/.env ./internal/.env

EXPOSE $GRPC_PORT

# Run the app
CMD ["./user-service", "-mode=deploy"]