# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download dependencies
COPY ../cmd/go.mod ../cmd/go.sum ./
RUN go mod download

# Copy the rest of the source code from cmd directory
COPY ../cmd/ ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built Go binary from the previous stage
COPY --from=builder /app/main .

# Copy the .env file
COPY ../cmd/.env .

# Expose the port the app runs on (change if needed)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
