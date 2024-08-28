# Use the official Golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire API directory to the working directory
COPY . .

# Build the API server
RUN go build -o api-server ./cmd

# Expose the port the API server listens on
EXPOSE 8080

# Run the API server
CMD ["./api-server"]
