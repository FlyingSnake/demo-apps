# Base image
FROM golang:1.16

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod .
RUN go mod download

# Copy the rest of the application
COPY . .

# Download dependencies (This will create go.sum file)
RUN go mod tidy

# Build the application
RUN go build -o main src/main.go

# Expose port 80
EXPOSE 80

# Command to run the application
CMD ["./main"]
