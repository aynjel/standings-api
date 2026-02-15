# Multi-stage build for optimal image size
FROM golang:1.21.6-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a static binary without cgo dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tabulation-api .

# Final stage - minimal runtime image
FROM alpine:latest

# Install runtime dependencies (ca-certificates for HTTPS)
RUN apk --no-cache add ca-certificates

# Create app directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/tabulation-api .

# Copy .env.example (optional, for reference)
COPY .env.example .

# Expose port
EXPOSE 4300

# Health check (optional)
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:4300/ || exit 1

# Run the application
CMD ["./tabulation-api"]
