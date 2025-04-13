# Build stage
FROM golang:1.24.1-alpine AS builder

# Set working directory
WORKDIR /app

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on

# Install build dependencies
RUN apk add --no-cache git ca-certificates && update-ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum* ./

# Download dependencies with BuildKit cache
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o mando -ldflags="-s -w" .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from the builder stage
COPY --from=builder /app/mando /app/mando

# Copy .env file - only if you want to include it in the container
# For production, consider using environment variables instead
COPY .env* ./

# Set the binary as executable
RUN chmod +x /app/mando

# Expose ports if needed
EXPOSE 8080

# Run the binary
CMD ["/app/mando"]