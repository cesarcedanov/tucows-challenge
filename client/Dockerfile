# Use the official Golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the entire Client directory to the working directory
COPY . .

# Build the CLI tool
RUN go build -o cli-client .

# Set the entry point to the CLI tool
ENTRYPOINT ["./cli-client"]
