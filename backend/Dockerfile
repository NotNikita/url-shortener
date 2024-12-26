# Multi-stage build
# Multi-stage builds allow you to separate the build environment from the runtime environment, 
# resulting in a much smaller final image.

# Stage 1: Build the app
FROM golang:1.23.3-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Cope the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. They'll be cached
RUN go mod download

# Build the App
RUN go build -o main .

# Stage 2: Create the final image
# It weight much smaller, cause doesnt include "golang" image and all the build tools
FROM alpine:latest

# Set the curr working dir
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY .env .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]