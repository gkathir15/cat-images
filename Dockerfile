# Start from a Debian-based image with the latest version of Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
ADD . /app

# Build the application inside the container
RUN go build -o main .

# Document that the service listens on port 8080
EXPOSE 8080

# Run the executable
CMD ["/app/main"]
