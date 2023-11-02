# syntax=docker/dockerfile:1

# Set the base image to the official Golang 1.21 image
FROM golang:1.21

# Set the working directory inside the container to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the container (dependency management)
COPY go.mod go.sum ./

# Download the Go module dependencies using 'go mod download'
RUN go mod download

# Copy all the files from the current directory into the container
COPY . .

# Build the Go application with CGO disabled for a static binary
# and target the Linux OS, then place the binary in /docker-goapp
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-goapp

# Expose port 8080 to allow external connections
EXPOSE 8080

# Define the command to run the application when the container starts
CMD ["/docker-goapp"]