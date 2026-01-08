# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy source code
COPY src/ ./src/

# Build the calculator
RUN cd src && go build -o /app/calculator main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/calculator /app/calculator

# Set the entrypoint to the calculator
ENTRYPOINT ["/app/calculator"]

# Default help message if no args provided
# CMD ["add", "0", "0"]
