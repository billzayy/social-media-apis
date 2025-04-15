# Use a minimal Golang base image
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy necessary files from the build context
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy the source code (cmd and internal folders)
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/

# Build the Go application
RUN GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

# Use a smaller base image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main /app/main

COPY --from=builder /app/internal /app/internal

# Command to run the application
CMD ["/app/main"]