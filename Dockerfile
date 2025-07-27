# Start from the official Golang image
FROM golang:1.24.1-alpine

# Set working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o bookservice

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./bookservice"]
