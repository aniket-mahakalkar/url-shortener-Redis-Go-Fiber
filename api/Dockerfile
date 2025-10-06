# Use official Golang image as build environment
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (if needed for go mod) and build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o url-shortener .

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/url-shortener .

# Expose port (change if your app uses a different port)
EXPOSE 8080

# Command to run the executable
CMD ["./url-shortener"]