# Use the official Golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire API directory to the working directory
COPY . .

RUN chmod 777 /app/cmd/cert/server.crt
RUN chmod 777 /app/cmd/cert/server.key

# Copy the cert directory to the /app/cmd/ directory in the container
COPY ./cmd/cert /app/cmd/cert

# Build the API server
RUN go build -o api-server ./cmd/main.go

# Expose the port the API server listens on
EXPOSE 8080

# Run the API server
CMD ["./api-server"]
