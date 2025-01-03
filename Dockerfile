# Stage 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Run
FROM debian:bullseye-slim

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main /app/

# Copy necessary directories (e.g., templates, static, and content)
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static
COPY --from=builder /app/content /app/content

# Install required runtime dependencies
RUN apt-get update && apt-get install -y libc6 && rm -rf /var/lib/apt/lists/*

# Expose the application port
EXPOSE 8989

# Command to run the application
CMD ["./main"]
