# Stage 1: Build the Go binary
FROM golang:1.24 as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o hello ./cmd/hello

# Stage 2: Create a minimal image
FROM gcr.io/distroless/base-debian11

WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/hello .

# Command to run the binary
CMD ["./hello"]