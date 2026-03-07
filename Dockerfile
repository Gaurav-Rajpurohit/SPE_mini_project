# Use official Go image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy project files
COPY . .

# Download dependencies
RUN go mod tidy

# Run tests
RUN go test ./...

# Build the application
RUN go build -o calculator

# Run the binary
CMD ["./calculator"]
